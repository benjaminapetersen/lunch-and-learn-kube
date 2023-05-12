# Jobs

## What are Jobs and CronJobs?

https://www.youtube.com/watch?v=OZAhYSDkhsI

- Three types of k8s jobs
  - Nonparallel jobs
    - usually a single pod
    - creates a new pod if the pod fails
    - no defined labels or label selectors
    - example:
      ```yaml
      apiVersion: batch/v1
      kind: Job
      metadata:
        name: hundred-fibonaccis
        # no labels needed
      spec:
        # no selector with label selectors, jobs will do this automatically
        template:
          spec:
            containers:
            - name: fibonaccis
              image: truek8s/hundred-fibonaccis:1.0
            # jobs are only designed to run pods which eventually terminate.
            # policy: OnFailure, Never are the only reasonable options.
            # if policy was Always, the job could never "finish"
            restartPolicy: OnFailure
        # specify the maximum number of times the job should restart the pod
        # before giving up.
        # backoff limit does have an exponential backoff delay:
        #   10sec,20,40....6 minutes max
        # default is 6 retries
        backoffLimit: 3
      ```
  - Parallel jobs with a fixed completion count
  - Parallel jobs with a task queue
  - Pods for Jobs are not deleted when completed
    - So we can query for logs later


- CronJobs
    - Runs k8s jobs at specifically scheduled moments in time
    - Modelled after Unix chron utility
    - Example schedules:
      - 5 mins after the start of every hour:
        - 5 * * * *
      - 1 minute before end of day each day:
        - 59 23 * * *
      - 15th day of every month
        - * * 15 * *
      - every 3 minutes
        - */3 * * * *
    - example cronjob:
        ```yaml
        apiVersion: batch/v1beta1
        kind: CronJob
        metadata:
            name: example
        spec:
            # once a minute (roughly)
            schedule: "*/1 * * * *"
            # job objects this CronJob will create
            jobTemplate:
                spec:
                    template:
                        spec:
                            containers:
                            - name: example
                              image: busybox:1.28
                              args:
                              - /bin/bash
                              - -c
                              - date
                            # We could skip restart since this runs often
                            restartPolicy: OnFailure
        ```
