
  <head>
  <script  type="text/javascript" src="https://openapi.map.naver.com/openapi/v3/maps.js?ncpClientId=yyi55imhvp">
      import Geolocation from "svelte-geolocation";
    let coords = [];
  </script>
  </head>

<h1>Home</h1>
Welcome to my website

<body>
  <!--<Geolocation getPosition bind:coords />-->

  <!--<div id="map" style="width:100%;height:400px;"></div>-->
   <div id="defaultmap" style="width:100%;height:400px;"></div>
  <!-- 
    Lets use Map API, find closest store or Restaurant 
  <pre>{JSON.stringify(coords)}</pre>
  <pre> First:  {coords[0]}, Second :{coords[1]} </pre>
  -->
  <script>

/*
     var mapOptions = {
     center: new naver.maps.LatLng(coords[1], coords[0]),
     zoom: 10
 };
  
   var map = new naver.maps.Map('map', mapOptions);
*/
   var mapDiv = document.getElementById('map'); // 'map'으로 선언해도 동일
   var defaultmap = new naver.maps.Map(mapDiv);


   </script>


</body>