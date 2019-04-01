use Ecomm;

var data = db.product.initializeUnorderedBulkOp();

data.insert({_id:1 , title : "Marvel's Spider Man(PS4)" , image :"https://m.media-amazon.com/images/I/51ndxZcdt+L._AC_UL872_FMwebp_QL65_.jpg", price:2690 ,rating:4.7 });
data.insert({_id:2 , title : "Red Dead Redemption - 2 (PS4)" , image :"https://m.media-amazon.com/images/I/51xDl3CQB1L._AC_UL872_FMwebp_QL65_.jpg", price:3890 ,rating:5.0 });
data.insert({_id:3 , title : "God of War - Standard Edition (PS4)" , image :"https://m.media-amazon.com/images/I/81uq-Xg1EWL._AC_UL872_FMwebp_QL65_.jpg", price:2284 ,rating:5.0 });
data.insert({_id:4 , title : "WWE 2K19 (PS4)",image:"https://m.media-amazon.com/images/I/915YrXCHZrL._AC_UL872_FMwebp_QL65_.jpg", price:1787 ,rating:4.6 });
data.insert({_id:5 , title : "Fifa 19 (PS4)",image:"https://m.media-amazon.com/images/I/71unWeOXWlL._AC_UL872_FMwebp_QL65_.jpg", price:2999 ,rating:4.8 });

data.execute();