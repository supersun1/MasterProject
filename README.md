# MasterProject
## Preface
This is a Master Degree Thesis/Project that were researched and implemented by **Lucero D., William F., Kostiantyn P., Dylan S., and Tommy Y.**

This project is advised by Dr. Wencen Wu of the San Jose State University

All rights reserved 2020.



## About
This Master Project focused on utilizing Natural Language Processing models to determine the accuracy and truthfulness of Covid-19 related news, articles, or texts.

From our model, we were able to achieve **97%** accuracy on every prediction.

This project's backend process is created using Google's [BERT](https://github.com/google-research/bert) and Facebook's [RoBERTa](https://pytorch.org/hub/pytorch_fairseq_roberta/) to accomplish this feat.

This accomplishment is explained at a high level in the following [video](https://www.youtube.com/watch?v=jdU5Zek-iJM)


## Front End Start Up Guide
Install node.js by downloading the latest here https://nodejs.org/en/download/ or using Homebrew through the terminal
node.js comes with npm 

Download the front-end code and open on your IDE of your choice.
On your terminal Navigate to the folder directory and run `npm start`

In the project directory, you can run:
### `npm start`

Runs the app in the development mode.\
Open [http://localhost:3000](http://localhost:3000) to view it in the browser.

The page will reload if you make edits.\
You will also see any lint errors in the console.

### `npm test`

Launches the test runner in the interactive watch mode.\
See the section about [running tests](https://facebook.github.io/create-react-app/docs/running-tests) for more information.

### `npm run build`

Builds the app for production to the `build` folder.\
It correctly bundles React in production mode and optimizes the build for the best performance.

The build is minified and the filenames include the hashes.\
Your app is ready to be deployed!

## Back End Start Up Guide
Our project is being hosted on Amazon Web Services. Through AWS we are utilizing a serverless architecture, which allows more flexibility and effectiveness when it comes to hosting overhead and prediction speeds. Our Front End is hosted on a publicly facing S3 bucket, which connects through AWS microservices (API Gateway & Lambda) to send requests and receive predictions from Sagemaker.

### Sagemaker Endpoints
For our project, we trained our ML models locally instead of hosting and training strictly on AWS. This was due to the cost and overhead that would be needed to accomplish our needs. In order to create our Sagemaker portion of the AWS architecture, we required Sagemaker Endpoints to host our models and create real-time predictions on requests. More information can be found in [AWS's Sagemaker Whitepaper](https://docs.aws.amazon.com/sagemaker/index.html).

We created Sagemaker Endpoints for both our BERT and RoBERTa models. In order to create these Sagemaker Endpoints, Sagemaker requires an inference script, pre-trained model, and model deployment notebook to be placed into a Notebook Instance within Sagemaker. The inference script is telling the pre-trained model how to load, receive input, predict, and push output. The model deployment notebook is used to package the inference script and pre-trained model, create a Sagemaker container and API, and create the Sagemaker Endpoint. Once the model deployment notebook runs and has completed the endpoint creation process, an S3 bucket is created for the Endpoint artifacts. This process needs to be completed for both the BERT and RoBERTa pre-trained models.


### Lambda & API Gateway
Our Lambda function is connecting the API Gateway to our Sagemaker Endpoints, as well as housing our classifier model. Through the utilization of Layers, we can use libraries and dependencies that require installation. The nature of Lambda does not allow the install of libraries as it is a temporary service that is only active when called upon. The Lambda function python script is used to calculate the final prediction accuracy, clean up the BERT and RoBERTa predictions, run the TF-IDF classifier, and invoke the Sagemaker endpoints.

The API Gateway is used to create a RESTful API which utilizes a POST method to send Requests to the Sagemaker Endpoint and return predictions to the Front End. To allow the S3 hosted Front End to call the API Gateway, we need to enable CORS. CORS (Cross-Origin Resource Sharing) allows services to access one another when in separate origins. CORS tends to not work out of the box, so you will need to manually enter the specified headers in the Method Response for each Method you have.    

### S3 Static Hosting
A publicly accessible S3 bucket is created for statically hosting our Front End. S3 provides hosting for the Front End in a static state, and dynamic functionality is provided by the microservices (API Gateway & Lambda). After the S3 Bucket is created and publicly accessible, import the Front End code into the bucket and use the provided bucket URL and append /index.html to the end. This will bring you to the Front-End splash page and begin to make requests to the Sagemaker Endpoints. Connecting the S3 bucket to AWS Cloudfront will allow your webpage to use an AWS provided SSL Certificate and Domain Name if you have one.



## Machine Learning Start Up Guide

### Model Training
For our project, we are using PyTorch as our framework. The NVIDIA’s CUDA is used to accelerate the training process. The pre-trained BERT and RoBERTa is imported through the Hugging Face Transformers library. We apply a label count of 2 for binary classification. The input CSV files are parsed and loaded into a pandas data frame, the true and false target set and combined into a single training data set. The training data is tokenized and converted by both the BERT and RoBERTa tokenizer. Next the training data set is passed to the models for training.  Once training is completed, the weights and biases is saved.

### Inference
For inference, the base BERT and RoBERTa models must be loaded and have the weights and biases loaded from our saved files. The input data is parsed and tokenized and converted to tensors. Then it is converted into a numpy array so we can use it for output.

