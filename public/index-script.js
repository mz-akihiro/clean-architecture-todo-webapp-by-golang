function addPost(){

    var jsonData = {
        Task: document.querySelector('#newtask input').value
    };
    console.log("post json data",jsonData)

    $.ajax({
        type : 'post',
        url : "http://localhost:8080/create-task",
        data : JSON.stringify(jsonData),
        contentType: 'application/JSON',
        scriptCharset: 'utf-8'
    })
    .then(
        function(data, textStatus, jqXHR){
            console.log(jqXHR.status)
            document.querySelector('#tasks').innerHTML += `
                <div class="task">
                    <span id="taskname">
                        ${document.querySelector('#newtask input').value}
                    </span>
                    <span id="taskId">${data.taskId}</span>
                    <button class="delete">
                        <i class="far fa-trash-alt"></i>
                    </button>
                </div>
            `;

            var current_tasks = document.querySelectorAll(".delete");   // 新旧問わず全ての削除ボタンにイベントリスナーを付与してるので非効率（改善項目）
            for(var i=0; i<current_tasks.length; i++){
                current_tasks[i].onclick = function(){
                    var deleteTask = {
                        TaskId: parseInt(this.parentNode.querySelector("#taskId").textContent)
                    };
                    var taskThis = this.parentNode; // thisの値を保存（ajax内だと指す値が変わるため）
                    $.ajax({
                        type : 'delete',
                        url : "http://localhost:8080/delete-task",
                        data : JSON.stringify(deleteTask),
                        contentType: 'application/JSON',
                        scriptCharset: 'utf-8'
                    })
                    .then(
                        function(data, textStatus, jqXHR){
                            console.log(jqXHR.status)
                            taskThis.remove()
                        },
                        function(jqXHR, textStatus, errorThrown){
                            console.log(jqXHR.status)
                            if (jqXHR.status >= 500) {
                                alert("server error")
                            }else if (jqXHR.status === 401){
                                alert("Token error, return to login page")
                                window.location.href="http://localhost:8080/login.html"
                            }else if (jqXHR.status >= 400) {
                                alert("request error")
                            }
                        }
                    )
                }
            }

            var tasks = document.querySelectorAll(".task");
            for(var i=0; i<tasks.length; i++){
                tasks[i].onclick = function(){
                    this.classList.toggle('completed');
                }
            }
            
            document.querySelector("#newtask input").value = "";
        },
        function(jqXHR, textStatus, errorThrown){
            console.log("status NO");
            console.log(jqXHR.status)
            if (jqXHR.status >= 500) {
                alert("server error")
            }else if (jqXHR.status === 401){
                alert("Token error, return to login page")
                window.location.href="http://localhost:8080/login.html"
            }else if (jqXHR.status >= 400) {
                alert("request error")
            }
        }
    );
}

////////////////////////////////////
document.querySelector('#push').onclick = function(){
    //alert("add Task")
    if(document.querySelector('#newtask input').value.length == 0){
        alert("Please Enter a Task")
    }
    else{
        addPost()
    }
}

function toggleMenu() {
    var menu = document.querySelector(".menu");
    if (menu.style.display === "none") {
      menu.style.display = "block";
    } else {
      menu.style.display = "none";
    }
}

function logout() {
    $.ajax({
        type : 'post',
        url : "http://localhost:8080/logout-api",
        contentType: 'application/JSON',
        scriptCharset: 'utf-8'
    })
    .then(
        function(data){
            console.log("delete status OK");
            window.location.href="http://localhost:8080/login.html"
        },
        function(jqXHR){
            console.log(jqXHR.status)
            console.log("delete status NO");
        }
    );
 }