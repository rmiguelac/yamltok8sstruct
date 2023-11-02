package main

import (
	"fmt"

	batchv1 "k8s.io/api/batch/v1"
	"k8s.io/client-go/kubernetes/scheme"
)

var jobdef = `
apiVersion: batch/v1
kind: Job
metadata:
  name: XXX
spec:
  ttlSecondsAfterFinished: 100
  template:
    spec:
      nodeName: nodeinpi
      containers:
      - name: XXX
        image: repo:5000/image:0.1-alpha
        args: ["YYY", "ZZZ"]
        volumeMounts:
        - mountPath: /path/to/folder
          name: recordings
      volumes:
      - name: recordings
        hostPath:
          path: /path/to/folder
          type: Directory
      restartPolicy: Never
  backoffLimit: 4
`

func main() {
	// read file from filesystem
	// decode to struct
	d := scheme.Codecs.UniversalDeserializer().Decode

	obj, _, err := d([]byte(jobdef), nil, nil)
	if err != nil {
		fmt.Printf("%#v", err)
	}

	job := obj.(*batchv1.Job)
	fmt.Printf("%#v\n", job)

}
