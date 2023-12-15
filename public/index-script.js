function addPost(){

    var jsonData = {
        Task: document.querySelector('#newtodo input').value
    };
    console.log("post json data",jsonData)

    $.ajax({
        type : 'post',
        url : "http://localhost:8080/create-todo",
        data : JSON.stringify(jsonData),
        contentType: 'application/JSON',
        scriptCharset: 'utf-8'
    })
    .then(
        function(data, textStatus, jqXHR){
            console.log(jqXHR.status)
            let dialogId = `ex-dialog-${data.todoId}`;
            document.querySelector('#todos').innerHTML += `
                    <div class="todo">
                        <span id="todoname">
                            ${document.querySelector('#newtodo input').value}
                        </span>
                        <div class="todoButton">
                            <button type="button" onclick="document.getElementById('${dialogId}').showModal()">更新</button>
                            <dialog id="${dialogId}" aria-labelledby="${dialogId}-title"> <!--dialog要素は名前を変えないと同じものを使い回す-->
                                <form method="dialog">
                                    <h3>todoを更新</h3>
                                    <div id=${dialogId}> <!--inputの値を取得するためにdiv要素を配置-->
                                        <p><label>更新名: <input type="text" placeholder="${document.querySelector('#newtodo input').value}"></label></p>
                                    </div>
                                    <div class="dialogButton">
                                        <span id="dialogId">${dialogId}</span>
                                        <span id="updateTodoId">${data.todoId}</span>
                                        <button type="button" onclick="this.closest('dialog').close();">キャンセル</button>
                                        <button class="todoChange" type="submit">更新</button>
                                    </div>
                                </form>
                            </dialog>
                            <button class="delete">
                                <span id="todoId">${data.todoId}</span>
                                <i class="far fa-trash-alt"></i>
                            </button>
                        </div>
                    </div>
                `;

            var current_tasks = document.querySelectorAll(".delete");   // 新旧問わず全ての削除ボタンにイベントリスナーを付与してるので非効率（改善項目）
            for(var i=0; i<current_tasks.length; i++){
                current_tasks[i].onclick = function(){
                    var deleteTask = {
                        TaskId: parseInt(this.parentNode.querySelector("#todoId").textContent)
                    };
                    var taskThis = this.closest(".todo");   // thisの値を保存（ajax内だと指す値が変わるため）
                    $.ajax({                                // なおdialog追加によりthisの指し示す値が<div class="todoButton">になったため
                        type : 'delete',                    // [this.closet]を使用して<div class="todo">を指定できるように変更(this.closetは親要素を探索する)
                        url : "http://localhost:8080/delete-todo",
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

            var current_tasks = document.querySelectorAll(".todoChange");
            for(var i=0; i<current_tasks.length; i++){
                current_tasks[i].onclick = function(){
                    console.log("更新dialog番号 : " + this.parentNode.querySelector("#dialogId").textContent)
                    console.log("更新内容 : " + document.querySelector('#'+this.parentNode.querySelector("#dialogId").textContent+' input').value)
                    console.log("todoId : " + this.parentNode.querySelector("#updateTodoId").textContent)
                    var UpdateTodo = {
                        TodoId: parseInt(this.parentNode.querySelector("#updateTodoId").textContent),
                        Todo: document.querySelector('#'+this.parentNode.querySelector("#dialogId").textContent+' input').value
                    };
                    $.ajax({
                        type : 'put',
                        url : "http://localhost:8080/update-todo",
                        data : JSON.stringify(UpdateTodo),
                        contentType: 'application/JSON',
                        scriptCharset: 'utf-8'
                    })
                    .then(
                        function(data, textStatus, jqXHR){
                            console.log(jqXHR.status)
                            window.location.href="http://localhost:8080/index.html";
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
                    );
                }
            }

            // 押された際に文字状に横線を引く
            /* var tasks = document.querySelectorAll(".todo");
            for(var i=0; i<tasks.length; i++){
                tasks[i].onclick = function(){
                    this.classList.toggle('completed');
                }
            } */
            document.querySelector("#newtodo input").value = "";
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
    if(document.querySelector('#newtodo input').value.length == 0){
        alert("Please Enter a Todo")
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