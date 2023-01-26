<script>

//import { navigate } from "svelte-navigator";
import { get_backend_url } from "../../components/Util.svelte";

    let loginInfo = {
    userid: "",
    password: "",

};

async function handleOnSubmit() {
  var res_status ='';
  var res_text='' ;
  var token='';
  var message;
  try {
        //console.log('Before send message')
        const controller = new AbortController();
        // 5 second timeout:

        const timeout = 5000 ;
        const timeoutId = setTimeout(() => controller.abort() , timeout);

        console.log(JSON.stringify(loginInfo));
        const url = get_backend_url();
        const res = await fetch(`${url}/login`, {
        method: 'POST',
        timeout : timeout,
        signal : controller.signal,
        body: JSON.stringify(loginInfo)
        });

        res_status = res.status;
        res_text =res.statusText;
        clearTimeout(timeoutId);

        const json_res = await res.json();
        token = json_res.token;
        message= json_res.message;
        console.log("After await json function : " , json_res);
      } catch (err){
        console.log("Fetch Error Msg: ", err);
        alert('Unexpected Error While response to Json', err.message);
      }


    if (res_status == 200){
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
  