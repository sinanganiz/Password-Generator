let copied = () =>{
    const el = document.createElement('textarea');
    el.value = document.querySelector("#randomPass").textContent;
    document.body.appendChild(el);
    el.select();
    document.execCommand('copy');
    document.body.removeChild(el);
    M.toast({html: 'Copied'})
  
}


let copy_mypass = (id) =>{
  const el = document.createElement('textarea');
  el.value = document.querySelector("#pass"+id).textContent;
  document.body.appendChild(el);
  el.select();
  document.execCommand('copy');
  document.body.removeChild(el);
  M.toast({html: 'Copied'})

}

 // Or with jQuery

 $(document).ready(function(){
    $('.modal').modal();
  });
       


