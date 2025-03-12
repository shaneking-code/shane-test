Steps to Follow:

CONFIGURE GCLOUD
1. Run `gcloud init` to set up a GCloud environment locally.
2. Authorize using `gcloud auth application-default login`.
3. Enable the storage API with:
   ```
   gcloud services enable storage.googleapis.com
   ```
4. Setup the service account with:
   ```
   gcloud iam service-accounts create NAME
   ```
5. Retrieve the associated email with:
   ```
   gcloud iam service-accounts list
   ```
6. Get the bucket name using:
   ```
   gcloud storage buckets list
   ```
7. Provision the service account for read-only access using:
   ```
   gcloud storage buckets add-iam-policy-binding BUCKET_URL \
     --member='serviceAccount:SERVICE_ACCOUNT_EMAIL' \
     --role='roles/storage.objectViewer'
   ```
8. Expose the service account credentials by setting the `GOOGLE_APPLICATION_CREDENTIALS`
   environment variable to `path/to/service-account-key.json`.

SET UP RESOURCES WITH TERRAFORM
1. Initialize Terraform using:
   ```
   terraform init
   ```
2. Define Terraform configuration, assigning and configuring providers (`google`, `random`).
3. Define resources to be created (bucket, bucket object).
4. Check formatting and validate with:
   ```
   terraform fmt
   terraform validate
   ```
5. Preview changes with:
   ```
   terraform plan
   ```
6. Apply the changes using:
   ```
   terraform apply
   ```

VERIFY WITH GCLOUD
1. Verify the bucket creation with:
   ```
   gcloud storage buckets list
   ```
2. Verify that the file is in the bucket using:
   ```
   gcloud storage ls <url>
   ```
3. Verify the contents of the file using:
   ```
   gcloud storage cat <url/filename>
   ```

GO SCRIPT
1. Import the requisite packages for Google Storage using:
   ```
   go get PKG_URL
   ```
2. Create a function to read a file object given the context, a GCloud client, a bucket, and the object itself.
3. In the main function:
   - Set up context, the client, and the bucket.
   - Iterate through bucket contents using the iterator API.
   - Read each object (in this case, only one) and output its attributes (e.g., name) and its content by calling the file read function.
4. Ensure the file and reader are closed using `defer`, and handle all errors properly.

