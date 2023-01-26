<script>
    import { Router, Route, Link } from "svelte-navigator";
    import { onMount } from 'svelte';
    import { get_backend_url } from "../../components/Util.svelte";

    //export let UserList;


	let posts = [];

	onMount(async () => {
    const url = get_backend_url()
		const res = await fetch(`${url}/post`);
		posts = await res.json();
	});
  </script>
  
  <div class="Posts">
    <h2>PostList</h2>
  
    <span>You're at post/list</span>
    <section class="posts">
        <ul>

        {#each posts as post}
            <li> Subject :
            {post.subject} 
            </li>
            <li> Writer : 
            {post.writer} 
            </li>
            <li> Content : 
            {@html post.content} 
            </li>
            
        {/each}
      </ul>
    </section>
    <nav>
      <Link to="../">Go back to the landing page</Link>
    </nav>

  </div>
  