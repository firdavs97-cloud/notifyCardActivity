## Summary
Write notification service to notify user for card activities on his card. You need to write HTTP server which accepts events in JSON format via POST method and stores them in some sort of storage 
(could store either in db or some basic memory storage). You also need to write worker/job which will notify those events to client. Notification can be mocked just by printing them to terminal.

- Tried to use only standard library
- Provided instructions on how to build and run the code
- Kept your solution simple
- Wrote some tests too

### Installation
Run Mysql service and create database. Use the sql defined in table.sql to create the required table.
Change the config file in config/default.yaml file with port and database full address in the form "root:yourpassword@tcp(localhost:3306)/dbname"
Run the command on the project root directory
```cmd
go run cmd/app/main.go
```

You can test the server by making post request with json body as given in the example below.

### Sample Events
```javascript
{
  "orderType": "Purchase",
  "sessionId": "29827525-06c9-4b1e-9d9b-7c4584e82f56",
  "card": "4433**1409",
  "eventDate": "2023-01-04 13:44:52.835626 +00:00",
  "websiteUrl": "https://amazon.com"
},

{
  "orderType": "CardVerify",
  "sessionId": "500cf308-e666-4639-aa9f-f6376015d1b4",
  "card": "4433**1409",
  "eventDate": "2023-04-07 05:29:54.362216 +00:00",
  "websiteUrl": "https://adidas.com"
},

{
  "orderType": "SendOtp",
  "sessionId": "500cf308-e666-4639-aa9f-f6376015d1b4", 
  "card": "4433**1409", 
  "eventDate": "2023-04-06 22:52:34.930150 +00:00",
  "websiteUrl": "https://somon.tj"
}
```
