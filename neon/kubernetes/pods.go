package kubernetes

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/tgs266/neon/neon/api"
	"github.com/tgs266/neon/neon/errors"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/kubernetes"
)

type podRepo struct {
	cs        *kubernetes.Clientset
	namespace string
	c         *gin.Context
}

func (r podRepo) ListByInstanceLabel(label string) []v1.Pod {
	res, err := r.cs.CoreV1().Pods(r.namespace).List(context.TODO(), metav1.ListOptions{
		LabelSelector: labels.Set(metav1.LabelSelector{MatchLabels: map[string]string{"app.kubernetes.io/instance": label}}.MatchLabels).String(),
	})
	if err != nil {
		errors.NewNotFound("pods not found", nil).Abort(r.c)
	}
	return res.Items
}

func (r podRepo) GetPod(name string) *v1.Pod {
	res, err := r.cs.CoreV1().Pods(r.namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		errors.NewNotFound("pods not found", nil).Abort(r.c)
	}
	return res
}

// getPodStatus returns status string calculated based on the same logic as kubectl
// Base code: https://github.com/kubernetes/kubernetes/blob/master/pkg/printers/internalversion/printers.go#L734
func (r podRepo) GetPodStatus(name string) api.PodStatus {
	pod := r.GetPod(name)
	restarts := 0
	readyContainers := 0

	reason := string(pod.Status.Phase)
	message := string(pod.Status.Message)
	if pod.Status.Reason != "" {
		reason = pod.Status.Reason
		message = pod.Status.Message
	}

	initializing := false
	for i := range pod.Status.InitContainerStatuses {
		container := pod.Status.InitContainerStatuses[i]
		restarts += int(container.RestartCount)
		switch {
		case container.State.Terminated != nil && container.State.Terminated.ExitCode == 0:
			continue
		case container.State.Terminated != nil:
			// initialization is failed
			if len(container.State.Terminated.Reason) == 0 {
				if container.State.Terminated.Signal != 0 {
					reason = fmt.Sprintf("Init: Signal %d", container.State.Terminated.Signal)
					message = container.State.Terminated.Message
				} else {
					reason = fmt.Sprintf("Init: ExitCode %d", container.State.Terminated.ExitCode)
					message = container.State.Terminated.Message
				}
			} else {
				reason = "Init:" + container.State.Terminated.Reason
				message = container.State.Terminated.Message
			}
			initializing = true
		case container.State.Waiting != nil && len(container.State.Waiting.Reason) > 0 && container.State.Waiting.Reason != "PodInitializing":
			reason = fmt.Sprintf("Init: %s", container.State.Waiting.Reason)
			message = container.State.Waiting.Message
			initializing = true
		default:
			reason = fmt.Sprintf("Init: %d/%d", i, len(pod.Spec.InitContainers))
			message = ""
			initializing = true
		}
		break
	}
	if !initializing {
		restarts = 0
		hasRunning := false
		for i := len(pod.Status.ContainerStatuses) - 1; i >= 0; i-- {
			container := pod.Status.ContainerStatuses[i]

			restarts += int(container.RestartCount)
			if container.State.Waiting != nil && container.State.Waiting.Reason != "" {
				reason = container.State.Waiting.Reason
				message = container.State.Waiting.Message
			} else if container.State.Terminated != nil && container.State.Terminated.Reason != "" {
				reason = container.State.Terminated.Reason
				message = container.State.Terminated.Message
			} else if container.State.Terminated != nil && container.State.Terminated.Reason == "" {
				if container.State.Terminated.Signal != 0 {
					reason = fmt.Sprintf("Signal: %d", container.State.Terminated.Signal)
					message = container.State.Terminated.Message
				} else {
					reason = fmt.Sprintf("ExitCode: %d", container.State.Terminated.ExitCode)
					message = container.State.Terminated.Message
				}
			} else if container.Ready && container.State.Running != nil {
				hasRunning = true
				readyContainers++
			}
		}

		// change pod status back to "Running" if there is at least one container still reporting as "Running" status
		if reason == "Completed" && hasRunning {
			if hasPodReadyCondition(pod.Status.Conditions) {
				reason = string(v1.PodRunning)
				message = pod.Status.Message
			} else {
				reason = "NotReady"
				message = pod.Status.Message
			}
		}
	}

	if pod.DeletionTimestamp != nil && pod.Status.Reason == "NodeLost" {
		reason = string(v1.PodUnknown)
		message = "Unknown"
	} else if pod.DeletionTimestamp != nil {
		reason = "Terminating"
		message = "Terminating"
	}

	if len(reason) == 0 {
		reason = string(v1.PodUnknown)
		message = "Unknown"
	}

	return api.PodStatus{
		Message: message,
		Status:  reason,
	}

}

func hasPodReadyCondition(conditions []v1.PodCondition) bool {
	for _, condition := range conditions {
		if condition.Type == v1.PodReady && condition.Status == v1.ConditionTrue {
			return true
		}
	}
	return false
}
