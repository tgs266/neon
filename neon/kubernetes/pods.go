package kubernetes

import (
	"context"

	"github.com/gin-gonic/gin"
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
