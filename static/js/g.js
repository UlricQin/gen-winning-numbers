/* common function */
function chk_msg(msg, f) {
    if (msg.length == 0) {
        $.layer({
            title: false,
            closeBtn: false,
            time: 1,
            dialog: {
                msg: '操作成功：）',
                type: 1
            },
            end: f
        });
    } else {
        $.layer({
            title: '出错了：（',
            dialog: {
                type: 3,
                msg: msg
            }
        });
    }
}

function chk_msg_and_reload(msg) {
    chk_msg(msg, function () {
        window.location.reload();
    });
}

function err_msg(msg, f) {
    if (msg.length > 0) {
        $.layer({
            title: false,
            closeBtn: false,
            time: 1,
            dialog: {
                msg: msg
            },
            end: f
        });
    }
}

function ok_msg(msg, f) {
    if (msg.length == 0) {
        msg = '操作成功：）'
    }
    $.layer({
        title: false,
        closeBtn: false,
        time: 1,
        dialog: {
            msg: msg,
            type: 1
        },
        end: f
    });
}

/* business function */
function reset_cfg() {
    var base_dir = $("#base_dir").val();
    if (base_dir.length == 0) {
        err_msg('全局目录没有配置呢……', function () {
            $("#base_dir").focus();
        });
    } else {
        window.location.href = "/admin/config/reset?dir=" + base_dir;
    }
}

function save_cfg() {
    $.get('/admin/config/sync', {}, function (msg) {
        chk_msg(msg);
    });
}

function load_ji_yi_list_of(duo_yuan) {
    $("#jyl").load('/admin/config/jyl/dy/' + duo_yuan);
    $("#ddl").empty();
}

function load_da_di_list_of(duo_yuan, ji_yi) {
    $("#ddl").load('/admin/config/ddl/dy/' + duo_yuan + '/jy/' + ji_yi);
}

function ji_yi_checked(dy_dir, jy_dir, is_checked) {
    var checked = is_checked ? '1' : '0';
    $.get('/admin/config/update/jyck', {'checked': checked, 'dy': dy_dir, 'jy': jy_dir}, function (msg) {
        err_msg(msg);
    });
}

function da_di_checked(duo_yuan, ji_yi, dd_name, is_checked) {
    var checked = is_checked ? '1' : '0';
    $.get('/admin/config/update/ddck', {'checked': checked, 'dy': duo_yuan, 'jy': ji_yi, 'name': dd_name}, function (msg) {
        err_msg(msg);
    });
}

function update_search_num(duo_yuan, ji_yi, dd_name, s_num) {
    $.get('/admin/config/update/snum', {'num': s_num, 'dy': duo_yuan, 'jy': ji_yi, 'name': dd_name}, function (msg) {
        err_msg(msg);
    });
}

function update_ji_yi_rong_cuo(type, val, dy) {
    $.get('/admin/config/update/jyrc', {'val': val, 't': type, 'dy': dy}, function (msg) {
        err_msg(msg);
    });
}

function update_da_di_rong_cuo(type, val, jy, dy) {
    $.get('/admin/config/update/ddrc', {'val': val, 't': type, 'jy': jy, 'dy': dy}, function (msg) {
        err_msg(msg);
    });
}

//function new_da_di() {
//    var dy = $("#tit_dy").text();
//    var jy = $("#tit_jy").text();
//    $.post('/admin/config/dd/new', {'dy': dy, 'jy': jy, 'name': $("#new_dd_name").val(), 's_num': $("#new_dd_search_num").val()}, function(msg) {
//        chk_msg(msg, function() {
//            load_da_di_list_of(dy, jy);
//        });
//    })
//}

function new_duo_yuan() {
    var min = $("#scope_min").val();
    var max = $("#scope_max").val();
    $("#new_duo_yuan_btn").attr('disabled', true);
    $.post('/admin/config/dy/new', {'min': min, 'max': max}, function(msg) {
        if (msg.length > 0) {
            err_msg(msg);
        } else {
            ok_msg('请检查磁盘文件')
            $("#new_duo_yuan_btn").attr('disabled', false);
        }
    });
}

function update_winning_number(order, val) {
    $.get('/admin/config/update/winning-number', {'order': order, 'val': val}, function(msg) {
        err_msg(msg);
    });
}

function compute() {
    $.get('/admin/compute/model2', {}, function(data) {
        $("#result").val(data);
    });
}