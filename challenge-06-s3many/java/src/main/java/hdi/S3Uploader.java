package hdi;
import java.io.File;
import com.amazonaws.services.s3.AmazonS3;
import com.amazonaws.services.s3.AmazonS3ClientBuilder;
import com.amazonaws.services.s3.model.PutObjectRequest;

public class S3Uploader {
    private static final String BUCKET_NAME = System.getenv("BUCKET");
    private static String FILE_NAME = "readme-java";
    private static final int NUM_UPLOADS = 100;

    public static void main(String[] args) {
        AmazonS3 s3Client = AmazonS3ClientBuilder.defaultClient();

        File file = new File("readme.md");
        for (int i = 0; i < NUM_UPLOADS; i++) {
            String loop = Integer.toString(i);
            String finalFileName = FILE_NAME+loop+".md";
        
            s3Client.putObject(new PutObjectRequest(BUCKET_NAME, finalFileName, file));
            System.out.println("Uploaded " + finalFileName + " to S3 bucket: " + BUCKET_NAME);

        }
        s3Client.shutdown();
    }
}
