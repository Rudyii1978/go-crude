const postContainer = document.getElementById("posts_container")

const fetchBtn = document.getElementById("fetchBtn")
fetchBtn.addEventListener("click", fetchPosts)

// Attach the function to the form's submit event
const form = document.getElementById('form');
form.addEventListener('submit', function (event) {
   // event.preventDefault(); // Prevent the form from submitting normally
   submitForm();
});


function submitForm() {
   // Get form elements
   var form_title = document.getElementById('form_title');
   var form_body = document.getElementById('form_body');

   // Create an object to hold the form data
   var formData = {
      title: form_title.value,
      body: form_body.value
   };

   // Post the data to the API endpoint
   fetch('http://localhost:5000/posts', {
      method: 'POST',
      headers: {
         'Content-Type': 'application/json'
      },
      body: JSON.stringify(formData)
   })
      .then(response => response.json())
      .then(data => console.log(data))
      .catch((error) => {
         console.error('Error:', error);
      });
}

async function fetchPosts() {
   console.log("Loading posts...");
   let data
   try {
      const response = await fetch("http://localhost:5000/posts");
      data = await response.json();
      console.log("Posts => ", data.posts);
   } catch (error) {
      console.error('Error:', error);
   }

   loadPosts(data.posts)
}

function loadPosts(posts) {
   // posts.forEach((element, index) => {
   //    console.log(`Post ${index} => ${element.Title}`);
   // });
   posts.forEach(post => {
      let postDiv = document.createElement("div")
      let postTitle = document.createElement("h3")
      let postBody = document.createElement("p")

      // Set classes for css
      postDiv.classList.add("post")
      postTitle.classList.add("title")
      postBody.classList.add("body")
      // Add data to title and body
      postTitle.innerHTML = post.Title
      postBody.innerHTML = post.Body
      // Add title and body to post container
      postDiv.append(postTitle, postBody)
      // Append post to DOM
      postContainer.appendChild(postDiv)
   });
}
