<script>
	import { Router, Route, Link } from "svelte-navigator";
	import NavLink from "./components/NavLink.svelte";
	import Home from "./routes/Home.svelte";
	
	import UserList from "./routes/user/UserList.svelte";
	import SignUp from "./routes/user/SignUp.svelte";
	import Login from "./routes/user/Login.svelte";

	import PostWrite from "./routes/post/Write.svelte";
	import PostList from "./routes/post/List.svelte";
	let is_auth = !!sessionStorage.getItem('token');
	import { auth_logout } from "./components/Util.svelte";

	/*
	function logout() {
			sessionStorage.clear();
			is_auth = !sessionStorage.getItem('token');
			alert("Logout, Session would be clear. is_auth : ", is_auth);
	}
	*/
	console.log('is_auth', is_auth,', token : ', sessionStorage.getItem('token') );

</script>

<Router>
	<header>
		<NavLink to="/"><h1>My site</h1></NavLink>
  
	  <nav>
		
			<!--<Link to="/about">About</Link>-->
	
		<NavLink to="/post/Write">PostWrite</NavLink>
		<NavLink to="/post/List">PostList</NavLink>

		{#if is_auth}
			<!--<a on:click={logout} href="/">Logout</a>-->
			<NavLink to="/user/list">UserList</NavLink>
			<a on:click={auth_logout} href="/">Logout</a>
		{:else}
			<NavLink to="/user/signup">SignUp</NavLink>
			<NavLink to="/user/login">Login</NavLink>
		{/if}
	  </nav>
	</header>
	<main>
		<Route path="/user/list" component={UserList} />
		<Route path="/user/signup" component={SignUp}/>
		<Route path="/user/login" component={Login}/>

		<Route path="/post/Write" component={PostWrite}/>
		<Route path="/post/List" component={PostList}/>

		<Route path="/" component={Home}/>
		<Route path="*" component={Home}/>

	</main>

</Router>
<style>
	main {
		text-align: center;
		padding: 1em;
		max-width: 240px;
		margin: 0 auto;
	}

	h1 {
		color: #1100ff;
		text-transform: uppercase;
		font-size: 4em;
		font-weight: 100;
	}

	@media (min-width: 640px) {
		main {
			max-width: none;
		}
	}
</style>
