
const Home = () => {

    return (
        <div className="max-w-6xl mx-auto">
            <div className="grid grid-cols-3 md:grid-cols-4 md:gap-x-4 py-4 px-4 md:px-0">
                {/* priducts  */}
                <div className="grid grid-cols-1 col-span-2 md:grid-cols-3 md:col-span-3 md:gap-x-4 px-2">

                    <div className="flex flex-col product-item">
                        <div className="img">
                            <figure>
                                <img className="" src="https://images.unsplash.com/photo-1574180566232-aaad1b5b8450?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxzZWFyY2h8NTZ8fHQlMjBzaGlydHxlbnwwfHwwfHx8MA%3D%3D&auto=format&fit=crop&w=800&q=60" alt="shirt-1"/>
                            </figure>
                        </div>
                        <div className="py-4 px-3 flex flex-col">
                            <h3>White Shirt</h3>
                            <h4>$ 10.00</h4>
                        </div>
                        <div className="flex justify-center items-center py-3 add-to-cart">
                            <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" className="bi bi-cart-fill" viewBox="0 0 16 16"> <path d="M0 1.5A.5.5 0 0 1 .5 1H2a.5.5 0 0 1 .485.379L2.89 3H14.5a.5.5 0 0 1 .491.592l-1.5 8A.5.5 0 0 1 13 12H4a.5.5 0 0 1-.491-.408L2.01 3.607 1.61 2H.5a.5.5 0 0 1-.5-.5zM5 12a2 2 0 1 0 0 4 2 2 0 0 0 0-4zm7 0a2 2 0 1 0 0 4 2 2 0 0 0 0-4zm-7 1a1 1 0 1 1 0 2 1 1 0 0 1 0-2zm7 0a1 1 0 1 1 0 2 1 1 0 0 1 0-2z"/> </svg>
                            <p>Add to cart</p>
                        </div>
                    </div>

                    <div className="flex flex-col product-item">
                        <div className="img">
                            <figure>
                                <img className="" src="https://images.unsplash.com/photo-1571945153237-4929e783af4a?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=987&q=80" alt="shirt-1"/>
                            </figure>
                        </div>
                        <div className="py-4 px-3 flex flex-col">
                            <h3>White Shirt</h3>
                            <h4>$ 10.00</h4>
                        </div>
                        <div className="flex justify-center items-center add-to-cart py-3">
                            <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" className="bi bi-cart-fill" viewBox="0 0 16 16"> <path d="M0 1.5A.5.5 0 0 1 .5 1H2a.5.5 0 0 1 .485.379L2.89 3H14.5a.5.5 0 0 1 .491.592l-1.5 8A.5.5 0 0 1 13 12H4a.5.5 0 0 1-.491-.408L2.01 3.607 1.61 2H.5a.5.5 0 0 1-.5-.5zM5 12a2 2 0 1 0 0 4 2 2 0 0 0 0-4zm7 0a2 2 0 1 0 0 4 2 2 0 0 0 0-4zm-7 1a1 1 0 1 1 0 2 1 1 0 0 1 0-2zm7 0a1 1 0 1 1 0 2 1 1 0 0 1 0-2z"/> </svg>
                            <p>Add to cart</p>
                        </div>
                    </div>

                    <div className="flex flex-col product-item">
                        <div className="img">
                            <figure>
                                <img className="" src="https://images.unsplash.com/photo-1627225924765-552d49cf47ad?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=987&q=80" alt="shirt-1"/>
                            </figure>
                        </div>
                        <div className="py-4 px-3 flex flex-col">
                            <h3>White Shirt</h3>
                            <h4>$ 10.00</h4>
                        </div>
                        <div className="flex justify-center items-center add-to-cart py-3">
                            <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" className="bi bi-cart-fill" viewBox="0 0 16 16"> <path d="M0 1.5A.5.5 0 0 1 .5 1H2a.5.5 0 0 1 .485.379L2.89 3H14.5a.5.5 0 0 1 .491.592l-1.5 8A.5.5 0 0 1 13 12H4a.5.5 0 0 1-.491-.408L2.01 3.607 1.61 2H.5a.5.5 0 0 1-.5-.5zM5 12a2 2 0 1 0 0 4 2 2 0 0 0 0-4zm7 0a2 2 0 1 0 0 4 2 2 0 0 0 0-4zm-7 1a1 1 0 1 1 0 2 1 1 0 0 1 0-2zm7 0a1 1 0 1 1 0 2 1 1 0 0 1 0-2z"/> </svg>
                            <p>Add to cart</p>
                        </div>
                    </div>

                    <div className="flex flex-col product-item">
                        <div className="img">
                            <figure>
                                <img className="" src="https://images.unsplash.com/photo-1564859228273-274232fdb516?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=987&q=80" alt="shirt-1"/>
                            </figure>
                        </div>
                        <div className="py-4 px-3 flex flex-col">
                            <h3>White Shirt</h3>
                            <h4>$ 10.00</h4>
                        </div>
                        <div className="flex justify-center items-center add-to-cart py-3">
                            <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" className="bi bi-cart-fill" viewBox="0 0 16 16"> <path d="M0 1.5A.5.5 0 0 1 .5 1H2a.5.5 0 0 1 .485.379L2.89 3H14.5a.5.5 0 0 1 .491.592l-1.5 8A.5.5 0 0 1 13 12H4a.5.5 0 0 1-.491-.408L2.01 3.607 1.61 2H.5a.5.5 0 0 1-.5-.5zM5 12a2 2 0 1 0 0 4 2 2 0 0 0 0-4zm7 0a2 2 0 1 0 0 4 2 2 0 0 0 0-4zm-7 1a1 1 0 1 1 0 2 1 1 0 0 1 0-2zm7 0a1 1 0 1 1 0 2 1 1 0 0 1 0-2z"/> </svg>
                            <p>Add to cart</p>
                        </div>
                    </div>

                    <div className="flex flex-col product-item">
                        <div className="img">
                            <figure>
                                <img className="" src="https://images.unsplash.com/photo-1623113562225-694f6a2ee75e?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=987&q=80" alt="shirt-1"/>
                            </figure>
                        </div>
                        <div className="py-4 px-3 flex flex-col">
                            <h3>White Shirt</h3>
                            <h4>$ 10.00</h4>
                        </div>
                        <div className="flex justify-center items-center add-to-cart py-3">
                            <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" className="bi bi-cart-fill" viewBox="0 0 16 16"> <path d="M0 1.5A.5.5 0 0 1 .5 1H2a.5.5 0 0 1 .485.379L2.89 3H14.5a.5.5 0 0 1 .491.592l-1.5 8A.5.5 0 0 1 13 12H4a.5.5 0 0 1-.491-.408L2.01 3.607 1.61 2H.5a.5.5 0 0 1-.5-.5zM5 12a2 2 0 1 0 0 4 2 2 0 0 0 0-4zm7 0a2 2 0 1 0 0 4 2 2 0 0 0 0-4zm-7 1a1 1 0 1 1 0 2 1 1 0 0 1 0-2zm7 0a1 1 0 1 1 0 2 1 1 0 0 1 0-2z"/> </svg>
                            <p>Add to cart</p>
                        </div>
                    </div>

                    <div className="flex flex-col product-item">
                        <div className="img">
                            <figure>
                                <img className="" src="https://images.unsplash.com/photo-1613852348851-df1739db8201?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=987&q=80" alt="shirt-1"/>
                            </figure>
                        </div>
                        <div className="py-4 px-3 flex flex-col">
                            <h3>White Shirt</h3>
                            <h4>$ 10.00</h4>
                        </div>
                        <div className="flex justify-center items-center add-to-cart py-3">
                            <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" className="bi bi-cart-fill" viewBox="0 0 16 16"> <path d="M0 1.5A.5.5 0 0 1 .5 1H2a.5.5 0 0 1 .485.379L2.89 3H14.5a.5.5 0 0 1 .491.592l-1.5 8A.5.5 0 0 1 13 12H4a.5.5 0 0 1-.491-.408L2.01 3.607 1.61 2H.5a.5.5 0 0 1-.5-.5zM5 12a2 2 0 1 0 0 4 2 2 0 0 0 0-4zm7 0a2 2 0 1 0 0 4 2 2 0 0 0 0-4zm-7 1a1 1 0 1 1 0 2 1 1 0 0 1 0-2zm7 0a1 1 0 1 1 0 2 1 1 0 0 1 0-2z"/> </svg>
                            <p>Add to cart</p>
                        </div>
                    </div>
                </div>
                {/* cart */}
                <div className="">
                    <div className="product-item">
                        <div className="cart-panel">
                            <h3>Cart</h3>
                        </div>
                        <div className="px-3 py-3">
                            <p>your shopping cart is empty</p>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    )
}


export default Home