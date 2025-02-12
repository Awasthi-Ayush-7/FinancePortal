<img align="right" src="https://user-images.githubusercontent.com/68017211/152626463-dad19147-475b-464e-8718-d4a5137958ca.png">
<br clear="left"/>

# FinanceX
The project's goal is to provide a one-stop financial platform for controlling your costs. This portal allows you to link your bank account and conduct transactions with other accounts in order to settle your balances. The portal also has tools to help you manage your spending by offering a concise summary of your monthly costs and classifying them. This is a course project for CEN5035 - Software Engineering.

## Brief Description
* **Landing page** - It provides a summary of what the application does and all its features, along with contact information of the authors and rating and reviews from other users.
* **Secured sign up and login** - User has the option to create a new account for sign up into existing account.
* **Add bank account** - Lets you link your bank account to add money to your wallet.
* **Shared expenses** - Lets you keep a track of shared expenses and split accordingly among your friends.
* **Expenditure summary** - It gives you a brief idea about your overall expenses by classifying them into different categories like groceries, rent, bills and so on.

## Technologies Used
* Golang
* Angular 13
* PostgreSQL
* HTML
* CSS
* JavaScript
* TypeScript

## Team Members
* Tushar Ranjan (Frontend)
* Mehul Jhaver (Backend)
* Sankalp Pandey (Frontend)
* Ayush Awasthi (Backend)

## GitHub Username Mapping
* HappyDash -> Sankalp Pandey
* tusharranjan719 -> Tushar Ranjan
* mehuljhaver4 -> Mehul Jhaver
* Awasthi-Ayush-7 -> Ayush Awasthi

## How to run
**Dashboard**
* Navigate to folder ../FinancePortal/financex-ui-dashboard
* Enter "npm run start"

**Bill-Split**
* Navigate to folder ../FinancePortal/financex-Backend/dashboard/frontend
* Enter "npm run start"

**Server**
* Navigate to folder ../FinancePortal/financex-Backend
* Launch Postgres: "sudo -u postgres psql", and then run "creadb test_bill" Exit and then run: "psql -f database/init/setup.sql -d test_bill"
* To build enter "go build -o bill-split"
* To run enter "./bill-split"

## Sprint 1
**Tasks Completed** 

* **Landing Page** 
  * It is the homepage of FinanceX which consists of various tabs on navigation bar such as Home, About, Features, Team and Login.
  * The FinanceX logo appears in the Home section.
  * The About section tells you everything you need to know about FinanceX.
  * The Feature section contains information about all of the application's features and services.
  * Information reagrding the members of the team can be found in the Team section.
  * We can find the testimonial section, which includes reviews from other users, by scrolling down.
  * The login tab on the navigation bar allows users to access their respective dashboards.
 
 * **Sign up**
    * The sign up section helps new users to create an account using their email address.
    * It also checks for a valid email address while creating a new account.
 * **Login**
    * The login section helps already existing users to login in and get access to their respective dashboards.

 **Sprint 1: videos**
* Frontend : https://youtu.be/f3Grv9Wpkk4
* Backend : https://youtu.be/dPFdobWmGa4
  
    

## Sprint 2
**Tasks Completed** 

  * **Frontend Testing**
    * Frontend of the project is integrated with backend.
    * Frontend testing was carried out with the help of Cypress. 
    * Every component in the project has been subjected to Cypress testing. 
    * Additionally, unit tests for the front are included for more thorough testing.
  
  * **Backend Testing**
    * The backend testing has been done with Postman. The testing can be viewed in sprint 1 video.
    * Units tests for the backend are provided for each middleware, database and route functions of the backend.

 **API Documentation**
  * Api Documentaion has been done using Swagger and can be viewed with the below link.
  * https://app.swaggerhub.com/apis-docs/tusharranjan719/FinanceX/1.0.0


## Sprint 3

**Tasks Completed**
  
   * **Dashboard** 
      * Greets the user when logged in.
      * Includes four cards for the quick overview of users expenses.
      * Lets user know about their balance, the amount they owe and their share in the splitted bills.
      * Has quick to-do list, to let the user quickly write and remember the bills or transaction they need to settle.
      * Small table for users last 5 recent transactions.
      * Has graphs that overlay and summarises users Annual, Weekly and Daily Average Expenses.
      * Added additional test cases to cover most functionalaties.


   * **Sidebar**
      * User Profile: Every participant user can update its information and change their current information.
      * Transactions: Users can see their history of their recent transactions.
      * Notifications: Every user receives a notification, whenever they are involved in a transaction.

   * **Navbar**
      * Includes the search bar for any assistance user needs on the application.
      * Has the quick notification option to see the recent transactions.
      * Includes the user profile button to display the entire user information. 

 **Sprint 3: videos**
  * https://youtu.be/3wTg8Z39GMc
  

## Sprint 4

**Project Description**

FinanceX is a service that simplifies bill splitting with friends and family. There are several options for splitting your invoices and payments through the site. You are greeted with an overview of our website and an introduction to the people behind it. You may easily join our portal by just logging in or creating a profile. The interface is interactive, allowing you to adjust the amount of money owed and divided in real time for every trip you plan. You may see your daily, weekly, or monthly spending in graphical and tabular forms on the dashboard. We have a simple navigation bar where you may get alerts whenever an expense is modified. You may even look up the charges that you were reminded to pay with our search option. The simple sidebar allows you to browse your profile, change it, and monitor your latest transactions. We centralize all of your shared spending and “I owe you’s” so that everyone knows who owes what. FinanceX makes life easier, whether you're splitting a holiday with friends, splitting rent with roommates, or repaying someone for lunch. 

**Demo Video**


https://user-images.githubusercontent.com/68017211/164120437-d1011109-fab1-49a6-8814-beb3666541eb.mp4



**Tasks Completed**

 * Developed bill split feature that let user split bill with different people.
 * User can create a bill split group and add participants with which they want to split bills
 * Expense can be created with expense name, amount, and payer. The payer then chooses with whom they split bills
 * Bill split page can be accessed from dashboard's transaction page
 * The overall expenses are visible in dashboard under different cards such as : Total expense, You owe, and You get back.
 * All transactions can also be viewed under Transaction tab in the sidebar.

**Sprint 4: Videos**
  * Deliverables : https://www.youtube.com/watch?v=QAt5Vaahg9Y
  
  * Cypress Testcases video 

https://user-images.githubusercontent.com/68017211/164121272-6b198e84-4d57-47a6-a51f-10f02b1dff30.mp4


  * Backend Testing Video


https://user-images.githubusercontent.com/68017211/164121237-67d8791a-4d2d-4e8f-a84f-dc03ca17cbce.mp4



 **API Documentation**
  * Api Documentaion has been done using Swagger and can be viewed with the below link.
  * https://app.swaggerhub.com/apis-docs/tusharranjan719/FinanceX/1.0.0


 **Project Board**
* https://github.com/tusharranjan719/FinancePortal/projects/3












