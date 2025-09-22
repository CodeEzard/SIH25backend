from google.cloud import vision
import os

# Resolve credential file path relative to this script
CRED_PATH = os.path.join(os.path.dirname(__file__), "vision-api.json")
os.environ['GOOGLE_APPLICATION_CREDENTIALS'] = CRED_PATH

client = vision.ImageAnnotatorClient()
image = vision.Image()
image_url = "https://images.besttemplates.com/wp-content/uploads/2024/06/Certificate8.jpg"
image.source.image_uri = image_url

response = client.text_detection(image=image)
texts = response.text_annotations
print("Texts:")
# Print the full text if available
if texts:
    print(texts[0].description)
