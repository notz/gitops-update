# gitops-update

GitHub Action that updates a single key in another GitHub repository.

## Example:

You have a `deployment.yaml` file in a `myorg/app-env` repository that has below content:

Add this to GitHub acion:

```text
- name: GitOps Update
	uses: simplycubed/gitops-update@0.15
	with:
		filename: "path/to/deployment.yaml"
		key-value: "image:${{ secrets.REGISTRY_LOGIN_SERVER }}/sampleapp:${{ github.sha }},tag:deployment"
		github-deploy-key: ${{ secrets.GITOPS_SSH_PRIVATE_KEY }}
		github-org:  'myorg'
		github-repo:  'app-env'
		username: 'john'
		email: 'example@gmail.com'
```

