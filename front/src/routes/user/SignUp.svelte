<script>

import { navigate } from "svelte-navigator";

    let signupUser = {
    userid: "",
    name: "",
    email: "",
    password: "",
    confirmpassword:"",
    level: 0,
};

async function handleOnSubmit() {
    let status = true;
    status = validate(signupUser);
    console.log(status);
    console.log(status == true);
    if (status == true) {
        delete signupUser.confirmpassword;
        const res = await fetch('http://52.78.73.92:8000/user', {
        method: 'POST',
        body: JSON.stringify(signupUser)
      });
      console.log(res);
      if (res.status == 200){
        alert('Signup Successful');
        navigate ('/');  
      } else{
        alert('Signup Failed : ', res.status, res.statusText)
      }
    } 
		
	}

function validate() {  
  let result = true;
  console.log('User :', signupUser, ',user.email :' ,signupUser.email );
  if ( signupUser.confirmpassword !== signupUser.password ) {
    alert('Confirm Password is not matched with password')
    result = false;
  }
  return result;
}

</script>
    
  <h1>Sign up</h1>
  <form class="content" on:submit|preventDefault={handleOnSubmit}  autocomplete="off" >
    <label>UserId<input type="text" bind:value={signupUser.userid} placeholder='UserId' autocomplete="off"/>
    </label>
    <label>Name<input type="text" bind:value={signupUser.name} placeholder='Name' autocomplete="off"/>
    </label>
    <label>Email<input type="email" bind:value={signupUser.email} placeholder='Email' autocomplete="off"/></label>
    
    <label>Password<input type="password" bind:value={signupUser.password} placeholder='Password' autocomplete="new-password"/></label>
    <label>Confirm Password<input type="password" bind:value={signupUser.confirmpassword} placeholder='ConfirmPassword' autocomplete="new-password"/></label>

    <button type="submit">
      Submit
    </button>
  </form>
  <p>
    {JSON.stringify(signupUser, 0, 2)}</p>
  