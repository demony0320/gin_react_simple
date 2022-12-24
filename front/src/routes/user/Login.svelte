<script>

import { navigate } from "svelte-navigator";

    let loginInfo = {
    userid: "",
    password: "",

};

async function handleOnSubmit() {
        console.log('Before send message')
        console.log(JSON.stringify(loginInfo))
        const res = await fetch('http://52.78.73.92:8000/login', {
        method: 'POST',
        body: JSON.stringify(loginInfo)
        });
      //console.log(res);
      const res_status = res.status;
      const res_text =res.statusText;
      var message ='';
      var token='';
      try {
        const json_res = await res.json();
        console.log(json_res);
        message =json_res.message;
        token =json_res.token;
      } catch (error){
        alert('Unexpected Error While response to Json')
      }
      if (res.status == 200){
        alert('Login Successful');
        sessionStorage.setItem('token',token);
        console.log(sessionStorage.getItem('token'));

        window.location.replace('/');
      } else{
        const msg = 'Login Failed. status: ' + res_status +  ', status msg: ' + res_text + ', server msg: ' + message ;
        alert(msg);
      }
    
		
	}


</script>
    
  <h1>Log in </h1>
  <form class="content" on:submit|preventDefault={handleOnSubmit}  autocomplete="off" >
    <label>UserId<input type="text" bind:value={loginInfo.userid} placeholder='UserId' autocomplete="off"/>
    </label>

    <label>Password<input type="password" bind:value={loginInfo.password} placeholder='Password' autocomplete="new-password"/></label>

    <button type="submit">
      Submit
    </button>
  </form>
  <p>
    {JSON.stringify(loginInfo, 0, 2)}</p>
  