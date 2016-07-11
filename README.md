## parking server by golang
---
### API
Function: Register

URL: http://ip:8080/v1/user/register

Method: POST

Body:

<pre><code>
 {
     "UserAccount": "18368493683",
     "UserPwd": "12345678",
     "CarLicense": "浙A12345"
 }
 </code></pre>

 Return:

<pre><code>
 {
     "Result": 0,
     "Err": "ok"
 }
</code></pre>


---
Function: Login

URL: http://ip:8080/v1/user/login

Method: POST

Body:


<pre><code>
 {
     "UserAccount": "18368493683",
     "UserPwd": "12345678"
 }
</code></pre>

 Return:

<pre><code>
 {
     "Result": 0,
     "Err": "ok",
     "Token": "519851815345456028818368493683",
     "CarLicenses": [
         {
             "CarLicense": "浙A12345"
         }
     ]
 }
</code></pre>


---
Function: Search ParkingLots by name

URL: http://ip:8080/v1/park/search

Method: GET

Requst:

http://ip:8080/v1/park/search?ParkName=name

 Return:

<pre><code>
 {
     "Result": 0,
     "Err": "ok",
     "Parks": [
         {
             "ParkName": "鄞州体育馆停车场",
             "ParkId": "3",
             "Longitude": 121.561854,
             "Latitude": 29.810474,
             "StoreyNum": 100,
             "EmptyNum": 50,
             "Money": 5,
             "ParkImg": "http://upload.admin5.com/2015/0728/1438073690416.jpg
         },
         {
             "ParkName": "鄞州公园停车场",
             "ParkId": "6",
             "Longitude": 121.546649,
             "Latitude": 29.811662,
             "StoreyNum": 100,
             "EmptyNum": 50,
             "Money": 5,
             "ParkImg": "http://upload.admin5.com/2015/0728/1438073690416.jpg
         },
         {
             "ParkName": "鄞州区政府停车场",
             "ParkId": "8",
             "Longitude": 121.546616,
             "Latitude": 29.817279,
             "StoreyNum": 100,
             "EmptyNum": 50,
             "Money": 5,
             "ParkImg": "http://upload.admin5.com/2015/0728/1438073690416.jpg
         }
     ]
 }
</code></pre>


---
Function: Show Nearby ParkingLots

URL: http://ip:8080/park/nearby

Method: POST

Body:


<pre><code>
 {
     "Longitude": 121.571636,
     "Latitude": 29.816461
 }
</code></pre>

Return :


<pre><code>
 {
     "Result": 0,
     "Err": "ok",
     "Parks": [
         {
             "ParkName": "鄞州体育馆停车场",
             "ParkId": "3",
             "Longitude": 121.561854,
             "Latitude": 29.810474,
             "StoreyNum": 100,
             "EmptyNum": 50,
             "Money": 5,
             "ParkImg": "http://upload.admin5.com/2015/0728/1438073690416.jpg
         },
         {
             "ParkName": "鄞州公园停车场",
             "ParkId": "6",
             "Longitude": 121.546649,
             "Latitude": 29.811662,
             "StoreyNum": 100,
             "EmptyNum": 50,
             "Money": 5,
             "ParkImg": "http://upload.admin5.com/2015/0728/1438073690416.jpg
         },
         {
             "ParkName": "鄞州区政府停车场",
             "ParkId": "8",
             "Longitude": 121.546616,
             "Latitude": 29.817279,
             "StoreyNum": 100,
             "EmptyNum": 50,
             "Money": 5,
             "ParkImg": "http://upload.admin5.com/2015/0728/1438073690416.jpg
         }
     ]
 }
 </code></pre>

 ---

 Function: BookParkingLots

 URL: http://ip:8080/v1/bookservices/book

 Method: POST

 Body:

 <pre><code>
  {
      "UserAccount": "18368493683",
      "CarLicense": "浙A12345",
      "Hours": 2.5,
      "ParkId": "1"
  }
  </code></pre>

  Return:

 <pre><code>
  {
      "Result": 0,
      "Err": "ok"
  }
 </code></pre>

 ---

 Function: GetParkInfoByParkId

 URL: http://ip:8080/v1/park/info

 Method: GET

 Requst:

 http://ip:8080/v1/park/info?ParkId=2

  Return:

 <pre><code>
 {
     "Result": 0,
     "Err": "ok",
     "StoreyNum": 100,
     "EmptyNum": 100,
     "Money": 5,
     "ParkImg": "http://upload.admin5.com/2015/0728/1438073690416.jpg"
}
 </code></pre>

 ---
Function: AddCarlicense

URL: http://ip:8080/v1/user/addcarlicense

Method: POST

Body:


<pre><code>
 {
     "UserAccount": "18368493683",
     "CarLicense": "浙B2233"
 }
</code></pre>

 Return:

<pre><code>
 {
     "Result": 0,
     "Err": "ok"
 }
</code></pre>
