

# Generated by: tyk-ci/wf-gen
# Generated on: Tuesday 08 February 2022 12:00:28 PM UTC

# Generation commands:
# ./pr.zsh -repos tyk -base master -branch tt-4363-el7 -title Sync from templates- el7 changes -p
# m4 -E -DxREPO=tyk


data "terraform_remote_state" "integration" {
  backend = "remote"

  config = {
    organization = "Tyk"
    workspaces = {
      name = "base-prod"
    }
  }
}

output "tyk" {
  value = data.terraform_remote_state.integration.outputs.tyk
  description = "ECR creds for tyk repo"
}

output "region" {
  value = data.terraform_remote_state.integration.outputs.region
  description = "Region in which the env is running"
}
