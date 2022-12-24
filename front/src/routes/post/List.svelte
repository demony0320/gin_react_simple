<script>
    import { Router, Route, Link } from "svelte-navigator";
    import { onMount } from 'svelte';

    //export let UserList;


	let posts = [];

	onMount(async () => {
		const res = await fetch(`http://52.78.73.92:8000/post`);
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
  