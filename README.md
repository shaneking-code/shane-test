# Steps
    ## CONFIGURE GCLOUD
    1. Run "gloud init" to set up a GCloud environment locally
    2. Authorize using "... auth application-default login"
    3. Allow usage of the storage API using "... services enable storage.googleapis.com"
    4. Setup service account using "gcloud iam service-accounts create NAME"
    5. Get the associated email with "... list"
    6. Get bucket name using "gcloud storage buckets list"
    7. Provision service account for read-only access using
       "gcloud storage buckets add-iam-policy-binding BUCKET_URL
       --member='serviceAccount:SERVICE_ACCOUNT_EMAIL'
       --role='roles/storage.objectViewer'"
    8. Expose the service account credentials by setting the GOOGLE_APPLICATION_CREDENTIALS
       environment variable to path/to/service-account-key.json

    ## SET UP RESOURCES WITH TERRAFORM
    1. Set up Terraform using "terraform init"
    2. Define Terraform config, assign and configure providers (google, random)
    3. Define resources to be created (bucket, bucket object)
    4. Check formatting and validate using "terraform fmt" and "terraform validate"
    5. Determine what will happen using "terraform plan"
    6. Apply using "terraform apply"

    ## VERIFY WITH GCLOUD
    1. Verify that the bucket was created using "gcloud storage buckets list"
    2. Verify that the file is in the bucket using "... ls <url>"
    3. Verify the contents of the file using "... cat <url/filename>"

    ## GO SCRIPT
    1. TODO

