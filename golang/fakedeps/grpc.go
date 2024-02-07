package fakedeps

import "k8s.io/client-go/kubernetes"

var _ = kubernetes.New