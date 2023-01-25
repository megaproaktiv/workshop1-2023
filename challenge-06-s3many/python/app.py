import boto3
import os

# Set up S3 client
s3 = boto3.client('s3')
bucket_name = os.environ['BUCKET']

# Define a function to upload an object
def upload_object(file_name):
    with open('readme.md', 'rb') as data:
        s3.upload_fileobj(data, bucket_name, file_name)
    print(f'Uploaded {file_name} to S3 bucket: {bucket_name}')


# Submit upload tasks to the thread pool

def main():
    for i in range(100):
        file_name = f'readme-python-{i}.md'
        upload_object(file_name)


if __name__ == "__main__":
    main()