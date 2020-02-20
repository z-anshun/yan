$(document).read(function () {
    DEFAULT_COOKIE_EXPIRE_TIME = 300;
    uname="";
    session="";
    uid=0;
    currentViedo=null;
    listedVideo=null;

    session=getCookie('session');
    uname=getCookie('user_name');
    //按钮的事件 注册

    initPage(function(){
        if(listedVideo!=null){
            currentViedo=listedVideo[0];
            selectVideo(listedVideo[0]['id'] );
        }
        $(".video-item").click(function(){

        });
        $(".del-video-button").click(function(){

        });
        $("#submit-comment").on("click",function() {

        })

    }) ;
    $('#regbtn').on('click',function (e) {

    });
    //登录
    $('#siginbtn').on('click',function (e) {

    });

    $('#siginnhref').on('click',function () {
        $('#regsubmit').hide();
        $('#siginsubmit').show();
    });

    $('#registerhref').on('click',function () {
        $('#regsubmit').show();
        $('#siginsubmit').hide();
    });

    $("#uploadfrom").on('submit',function (e) {

    });

    $('.close').on('click',function (e) {

    });

    $('#logout').on('click',function () {

    })
});

function initPage(callback) {

}

function setCookie(cname,cvalue,exmin) {

}

function getCookie(cname) {

}

function selectVideo(vid){

}

function refreshComments(vid){

}

function popupNotificationMsg(msg){

}
function popupErrorMsg(msg){

}

function htmlCommentListElement(cid,author,content){

}

function htmlVideoListElement(vid,name,ctime){

}

//异步
//user operations
function registerUser(callback){

}

function signinUser(callback ) {

}

function getUserId(callback){

}

//video operation
function createVideo(vname,callback) {

}

function listAllVideos(callback) {

}

function deleteVideo(vid,callback){


}

//comments operation
function postComment(vid,content,callback){

}

function listAllComments(vid,callback){

}