run-terraform:
	terraform -chdir=infrastructure/terraform init;
	terraform -chdir=infrastructure/terraform plan;
	terraform -chdir=infrastructure/terraform apply;
