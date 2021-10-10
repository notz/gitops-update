# gitops-update

GitHub Action that updates a single key in another GitHub repository.

## Examples:

You have a `deployment.yaml` file in a `myorg/app-env` repository that has below content:

Add this to GitHub acion:

```text
# example updating a single image
- name: GitOps Update
  uses: simplycubed/gitops-update@0.15
  with:
    filename: "path/to/deployment.yaml"
      key-value: "spec.template.spec.containers[0].image:gcr.io/<repo>/app:${{ secrets.COMMIT_SHA}}"
      github-deploy-key: ${{ secrets.GITOPS_SSH_PRIVATE_KEY }}
      github-org:  'myorg'
      github-repo:  'app-env'
      username: 'john'
      email: 'example@gmail.com'
```

```text
# example updating a container and an initContainer
- name: GitOps Update
  uses: simplycubed/gitops-update@0.15
  with:
    filename: "path/to/deployment.yaml"
      key-value: spec.template.spec.containers[0].image:gcr.io/<repo>/app:${{ secrets.COMMIT_SHA}},spec.template.spec.initContainers[0].image:gcr.io/<repo>/db-migrations:${{ secrets.COMMIT_SHA}}
      github-deploy-key: ${{ secrets.GITOPS_SSH_PRIVATE_KEY }}
      github-org:  'myorg'
      github-repo:  'app-env'
      username: 'john'
      email: 'example@gmail.com'
```

