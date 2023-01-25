const AWS = require('aws-sdk');

// Create an S3 client
const s3 = new AWS.S3();

// The name of the bucket and the file to upload
const bucket = process.env.BUCKET;
const file = 'readme.md';

// Upload the file n times
for (let i = 0; i < 100; i++) {
  let key =  "readme-node-"+i+".md"
  s3.upload({
    Bucket: bucket,
    Key: key,
    Body: require('fs').createReadStream(file)
  }, (err, data) => {
    if (err) {
      console.log(err);
    } else {
      console.log(`File uploaded successfully: `+key);
    }
  });
}
