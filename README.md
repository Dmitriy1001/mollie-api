#### 1. SignUp https://my.mollie.com/dashboard/signup?lang=en
#### 2. Fill some basic information about your company with fictitious data. It is recommended to select the country US or UK, as Germany requires a VAT number.
##### ![image](https://user-images.githubusercontent.com/48877778/182021544-12ab8dd2-dc21-4d66-8f1f-7d9c63b9e3c2.png)
You can stop at the "organization documentation" stage, the main thing is to choose payment methods, otherwise it will be impossible to create a payment
#### 3. Go to the API keys section and copy the Test API key.
##### ![image](https://user-images.githubusercontent.com/48877778/182022751-d933899c-2348-4e80-8229-9cd7f7967d7b.png)
#### 4. Create an environment variable MOLLIE_API_TOKEN with a test key to see how the examples work (uncomment them and substitute your payment id). If you will use the postman collection, set the variable test-api-key.  
##### Payments can be seen in the Transactions/Payments section, select the test mode in the upper right corner. ![image](https://user-images.githubusercontent.com/48877778/182023379-01122712-6b5c-4e40-87d5-bf1ed50300a0.png)

