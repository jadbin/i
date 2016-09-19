$(function (){
	// Code Highlight
	$("pre").each(function(i, block) {
		hljs.highlightBlock(block);
	});
});