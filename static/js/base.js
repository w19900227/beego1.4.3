// $.fn.serializeObject = function(){
//     var o = {};
//     var a = this.serializeArray();
//     $.each(a, function() {
//     	if (this.name.substr(-2) == "[]"){
//         	this.name = this.name.substr(0, this.name.length - 2);
//         	// o[this.name] = [];
//     	}
//         if (o[this.name]) {
//             if (!o[this.name].push) {
//                 o[this.name] = [o[this.name]];
//             }
//             o[this.name].push(this.value || '');
//         } else {
//             o[this.name] = this.value || '';
//         }
//     });
//     return o;
// };

$.fn.serializeObject = function(){
    var o = {};
    var a = this.serializeArray();
    var c = new Array()

    $.each(a, function() {
        if (this.name) {
            if (this.name.substr(-2) == "[]") {
                this.name = this.name.substr(0, this.name.length - 2);
                c.push(this.value || '');
                o[this.name] = c;
            } else {
                o[this.name] = this.value || '';
            }
        } else {
            o[this.name] = this.value || '';
        }
    });
    return o;
};

