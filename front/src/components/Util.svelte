
<script context="module">
  import { navigate } from "svelte-navigator";
  const pkgConfig = require('../../package.json');

    export async function auth_fetch(url) {
      console.log('auth_fetch start');
      try{
        let response = await fetch(url,
          { 
              headers : {'Authorization' :  `Bearer ${sessionStorage.getItem("token")}`} 
          }
        );
          let status =  await response.status; 
          let json = await response.json();
          if ( status != 200) {
            alert('logout , go to main')
            auth_logout();
          }
          return json;
        
        }catch(err){
            alert('logout , go to main')
            auth_logout();
      } 
    }

    export function auth_logout(){
      sessionStorage.clear();
			alert("Logout, Session would be clear");
      navigate ('/');  

    }
    export function get_backend_url() {
      alert(pkgConfig.backendUrl);
      console.log(`BackendUrl : ${pkgConfig.backendUrl}`);
      return pkgConfig.backendUrl;
    }


  export function imageUploader(image, callback) {
    console.log("Log : Image Uploader", image)
    var data = new FormData();
    data.append('image', image);

    var xhr = new XMLHttpRequest();
    xhr.open('POST', 'http://52.78.73.92:8000/upload', true);
    xhr.setRequestHeader('Authorization', `Bearer ${sessionStorage.getItem("token")}`);
    xhr.onreadystatechange = function() {
      if (xhr.readyState === 4) {
        var response = JSON.parse(xhr.responseText);
        if (response.status === 200 && response.success) {
          callback(response.data.link);
        } else {
          var reader = new FileReader();
          reader.onload = function(e) {
            callback(e.target.result);
          };
          reader.readAsDataURL(image);
        }
      }
    }
    xhr.send(data);
}
</script>

