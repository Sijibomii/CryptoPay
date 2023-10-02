import { Link } from "react-router-dom"

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
                            <h3>White Shirt packg I</h3>
                            <h4>$ 10.00</h4>
                        </div>
                        <div className="flex justify-center items-center py-3 add-to-cart cursor-pointer">
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
                            <h3>White Shirt packg II</h3>
                            <h4>$ 8.00</h4>
                        </div>
                        <div className="flex justify-center items-center add-to-cart cursor-pointer py-3">
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
                            <h3>White Shirt packg III</h3>
                            <h4>$ 4.00</h4>
                        </div>
                        <div className="flex justify-center items-center add-to-cart cursor-pointer py-3">
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
                            <h3>White Shirt packg IV</h3>
                            <h4>$ 6.00</h4>
                        </div>
                        <div className="flex justify-center items-center add-to-cart cursor-pointer py-3">
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
                            <h3>White Shirt V</h3>
                            <h4>$ 2.00</h4>
                        </div>
                        <div className="flex justify-center items-center add-to-cart cursor-pointer py-3">
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
                            <h3>White Shirt VI</h3>
                            <h4>$ 3.00</h4>
                        </div>
                        <div className="flex justify-center items-center add-to-cart cursor-pointer py-3">
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
                            {/* <p>your shopping cart is empty</p> */}
                            <div className="top mb-6 cart-products pb-3">
                                <div className="mb-2">
                                    <div className="flex items-center justify-between">
                                        <h6 className="text-sm">2 x White Shirt VI</h6>
                                        <h5>$ 6.00</h5>
                                    </div>
                                </div>

                                <div className="mb-2">
                                    <div className="flex items-center justify-between">
                                        <h6 className="text-sm">2 x White Shirt V</h6>
                                        <h5>$ 4.00</h5>
                                    </div>
                                </div>
                            </div>
                            
                            <div className="bottom mt-4">
                                <div className="">
                                    <div className="">
                                        <h3 className="text-xl font-bold w-full ml-auto mb-3">Total: $ 10.00</h3>
                                    </div>
                                </div>

                                <div className="flex items-center justify-between">
                                    <div className="flex items-center justify-between cart-border p-1">
                                        <svg xmlns="http://www.w3.org/2000/svg" width="12" height="12" fill="currentColor" className="bi bi-trash" viewBox="0 0 16 16"> <path d="M5.5 5.5A.5.5 0 0 1 6 6v6a.5.5 0 0 1-1 0V6a.5.5 0 0 1 .5-.5zm2.5 0a.5.5 0 0 1 .5.5v6a.5.5 0 0 1-1 0V6a.5.5 0 0 1 .5-.5zm3 .5a.5.5 0 0 0-1 0v6a.5.5 0 0 0 1 0V6z"/> <path fill-rule="evenodd" d="M14.5 3a1 1 0 0 1-1 1H13v9a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V4h-.5a1 1 0 0 1-1-1V2a1 1 0 0 1 1-1H6a1 1 0 0 1 1-1h2a1 1 0 0 1 1 1h3.5a1 1 0 0 1 1 1v1zM4.118 4 4 4.059V13a1 1 0 0 0 1 1h6a1 1 0 0 0 1-1V4.059L11.882 4H4.118zM2.5 3V2h11v1h-11z"/> </svg>
                                        <h5 className="ml-1 text-sm ">Empty cart</h5>
                                    </div>
 
                                    <div className="bg-green-500 p-1 rounded-[5px]">
                                        <Link className="text-sm text-white" to='/cart'>Checkout</Link>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    )
}


export default Home