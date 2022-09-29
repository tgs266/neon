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

type serviceRepo struct {
	cs        *kubernetes.Clientset
	namespace string
	c         *gin.Context
}

func (r serviceRepo) ListByInstanceLabel(label string) []v1.Service {
	res, err := r.cs.CoreV1().Services(r.namespace).List(context.TODO(), metav1.ListOptions{
		LabelSelector: labels.Set(metav1.LabelSelector{MatchLabels: map[string]string{"app.kubernetes.io/instance": label}}.MatchLabels).String(),
	})
	if err != nil {
		errors.NewNotFound("services not found", err).Panic()
	}
	return res.Items
}
