<script>
  import { onMount } from "svelte";
  import { navigate } from "svelte-navigator";
  import { object_without_properties } from "svelte/internal";
  import { imageUploader } from "../../components/Util.svelte";

  let editor;
	var quill;
	export let toolbarOptions = [
		[{ header: 1 }, { header: 2 }, "blockquote", "link", "image", "video"],
		["bold", "italic", "underline", "strike"],
		[{ list: "ordered" }, { list: "ordered" }],
		[{ align: [] }],
		["clean"]
	];

  let post_form_data = {
    category:"humor",
    writer: "jeffjang",
    subject: "",
    content :"",

};

const handleOnSubmit = async e => {
  e.preventDefault()
  //post_form_data.content = editor.children[0].innerHT`ML;
  
  var contents = quill.getContents();
  console.log(typeof(contents), contents)
  //console.log("root innerhtml :", quill.root.innerHTML)
  //var temp_editor = editor;
  post_form_data.content = editor.firstChild.innerHTML;
  
  var doc = document.querySelector(".ql-editor > p")
  var imgs = document.querySelectorAll(".ql-editor > p > img")
  //img.src="sleepy.jpg"
  //var new_doc = doc.removeChild(img)
  const formData = new FormData()

  
  for (let i = 0 ; i< imgs.length; i++){
    console.log("querySelectALl, data : ", typeof(imgs[i]) ,imgs[i].src.split(",")[1])
    let base_response = await fetch(imgs[i].src)
    let img_blob= await base_response.blob()
    doc.removeChild(imgs[i])
    const name = 'img' + i 
    formData.append(name,img_blob,name + '.jpg')
  }

  console.log(typeof(doc), typeof(imgs))
  //console.log('after remove child : ', new_doc)
  //var doc = new DOMParser().parseFromString(temp_editor.firstChild.innerHTML, "image/svg+xml");  
  console.log('After remove : ', doc)
  //doc = doc.querySelector('p')
  //console.log(doc, typeof(doc))

 // var images = doc.getElementsByTagName('img'); 

  //console.log("images!", images, typeof(images))
  //doc.removeChild(doc.querySelectorAll('img')[0]); 
  //console.log("after remove ", doc)
  
  //var texts = quill.getText();

  /*
  for (var op of contents.ops) 
  {
    //formData.append(op.insert)
    //console.log("type of : ", typeof(op.insert))
    if ( typeof(op.insert) == 'object') {
      //console.log("Object, ", typeof(op.insert))
      console.log("Object :" , op.insert);

    } else {
      //console.log("Not Object, ", typeof(op.insert))
    //console.log(op);
    console.log("String :" , op.insert);
    }
    //console.log("typeof op: ", typeof(op), " type of insert : ", typeof(op.insert))
  }
*/

  try {
    if (imgs.length > 0) {
      console.log('formData : ', formData.entries())
      const upload_resp = await fetch ('http://52.78.73.92:8000/upload' , {
        method: 'POST',
        headers: {
        'Accept': 'application/json',
        'Content-Type': 'multipart/form-data',
      },
        body: formData
      })
      alert('file upload successful', upload_resp)
      }
  }
  catch(err){
    console.log("Error : ", err ) 
    alert('Upload failed, try agian')
    abort()
  }


   const res = await fetch('http://52.78.73.92:8000/post', {
        method: 'POST',
        body: JSON.stringify(post_form_data)
      });
      console.log(res);
      if (res.status == 200){
        alert('Upload Post Successful');
        navigate ('/');  
      } else{
        alert('Upload Post Failed : ', res.status, res.statusText)
      }

}
/*
var IMGUR_CLIENT_ID = 'bcab3ce060640ba';
var IMGUR_API_URL = 'https://api.imgur.com/3/image';  

function imageHandler(image, callback) {
  console.log("image Handler");
  var data = new FormData();
  data.append('image', image);

  var xhr = new XMLHttpRequest();
  xhr.open('POST', IMGUR_API_URL, true);
  xhr.setRequestHeader('Authorization', 'Client-ID ' + IMGUR_CLIENT_ID);
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
*/
  function logFunction(){
    console.log("this is just log");
  }

  onMount(async () => {
		const { default: Quill } = await import("quill");
    //let quill = new Quill(editor, {
    quill = new Quill(editor, {
      modules: {
        toolbar: toolbarOptions,
      },
      theme: "snow",
      placeholder: "Write your story...",
      imageHandler: imageUploader // this never fucking work

    });
  });
</script>

<style>
  @import 'https://cdn.quilljs.com/1.3.6/quill.snow.css';
</style>

<form class="content" id="postWriteForm" on:submit|preventDefault={handleOnSubmit}  autocomplete="off" >
   
    <label>Subject<input type="text" bind:value={post_form_data.subject} placeholder='Subject' autocomplete="off"/>
    </label>
    <div bind:this={editor} />
  <!--<textarea name="text" style="display:none" id="hiddenArea" bind:value={post_form_data.context}></textarea>-->

  <button type="submit">
    Submit
  </button>
</form>

