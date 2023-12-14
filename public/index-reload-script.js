document.addEventListener("DOMContentLoaded", reloadcheckPost);
function reloadcheckPost(){
    console.log("reload")
    $.ajax({
        type : "post",
        url : "http://localhost:8080/read-todo",
    })
    .then(
        function(data){
            console.log("status OK");
            console.log(data)
            if (data != null){
                for (let i = 0; i < data.length; i++) {
                    let task = data[i];
                    console.log(i, " : ", task.TodoId + ": " + task.Todo);
                    reloadTodos(task.TodoId, task.Todo)
                }
            }
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

function reloadTodos(todoId, todo){
    let dialogId = `ex-dialog-${todoId}`;
    document.querySelector('#todos').innerHTML += `
            <div class="todo">
                <span id="todoname">
                    ${todo}
                </span>
                <div class="todoButton">
                    <button type="button" onclick="document.getElementById('${dialogId}').showModal()">更新</button>
                    <dialog id="${dialogId}" aria-labelledby="${dialogId}-title"> <!--dialog要素は名前を変えないと同じものを使い回す-->
                        <form method="dialog">
                            <h3>todoを更新</h3>
                            <div id=${dialogId}> <!--inputの値を取得するためにdiv要素を配置-->
                                <p><label>更新名: <input type="text" placeholder=${todo}></label></p>
                            </div>
                            <div class="dialogButton">
                                <span id="dialogId">${dialogId}</span>
                                <button type="button" onclick="this.closest('dialog').close();">キャンセル</button>
                                <button class="todoChange" type="submit">更新</button>
                            </div>
                        </form>
                    </dialog>
                    <button class="delete">
                        <span id="todoId">${todoId}</span>
                        <i class="far fa-trash-alt"></i>
                    </button>
                </div>
            </div>
        `;

    var current_tasks = document.querySelectorAll(".delete");
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
                    taskThis.remove();
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

    var current_tasks = document.querySelectorAll(".todoChange");
    for(var i=0; i<current_tasks.length; i++){
        current_tasks[i].onclick = function(){
            console.log("更新dialog番号 : " + this.parentNode.querySelector("#dialogId").textContent)
            console.log("更新内容" + document.querySelector('#'+this.parentNode.querySelector("#dialogId").textContent+' input').value)
            console.log("todoId : " + this.parentNode.querySelector("#updateTodoId").textContent)
        }
    }

    // 押された際に文字状に横線を引く
    /* var tasks = document.querySelectorAll(".todo");
    for(var i=0; i<tasks.length; i++){
        tasks[i].onclick = function(){
            this.classList.toggle('completed');
        }
    } */
}