// Access MongoDB's `newDB` database
const db = connect('mongodb://user:password@localhost:27017/newDB');

db.discs.drop();

db.discs.insertMany([
    {"name": "Crave", "img": "https://images-na.ssl-images-amazon.com/images/I/612lLSPhT1L.jpg", "price": "$20", "type": "Fairway"},
	{"name": "Destroyer", "img": "https://m.media-amazon.com/images/I/813mU5I8bFL.jpg", "price": "$20", "type": "Driver"},
	{"name": "Virus", "img": "https://images-na.ssl-images-amazon.com/images/I/51vEEQrB6yL.jpg", "price": "$20", "type": "Driver"},
	{"name": "Leopard3", "img": "https://m.media-amazon.com/images/I/81VBeQuv9wL.jpg", "price": "$18", "type": "Fairway"},
	{"name": "Resistor", "img": "https://m.media-amazon.com/images/I/51ygHUFkaqL.jpg", "price": "$18", "type": "Fairway"},
	{"name": "Hatchet", "img": "https://us.ftbpic.com/product-amz/westside-discs-origio-burst-hatchet-fairway-disc-golf-driver-great/51eJmz-1MTL._AC_SR480,480_.jpg", "price": "$18", "type": "Fairway"},
	{"name": "Maestro", "img": "https://www.discgolfmarket.com/cdn/shop/products/DM_ActivePremium_Maestro_1024x1024_bac409dd-62f1-42ce-85f0-56cb75b11cf2_512x512.jpg?v=1597850564", "price": "$18", "type": "Midrange"},
	{"name": "MX3", "img": "https://www.pbsports.com/cdn/shop/files/MX3-400Glow-1_400x_e4e0cf1d-e4f7-4068-b360-b8056ed75366.jpg?v=1724871751&width=533", "price": "$18", "type": "Midrange"},
	{"name": "Roc3", "img": "https://discgolffanatic.com/wp-content/uploads/2020/04/Innova-Champion-plastic-roc3-midrange.webp", "price": "$18", "type": "Midrange"},
	{"name": "P2", "img": "https://i0.wp.com/discgolffanatic.com/wp-content/uploads/2021/12/Discmania-P2-putter.jpg?resize=640,640&ssl=1", "price": "$18", "type": "Putter"},
	{"name": "Roach", "img": "https://i0.wp.com/discgolffanatic.com/wp-content/uploads/2021/11/Discraft-Roach-Putter.webp?resize=640,640&ssl=1", "price": "$18", "type": "Putter"},
	{"name": "Sensei", "img": "https://m.media-amazon.com/images/I/61B6634eizL.jpg", "price": "$18", "type": "Putter"}
  ]);