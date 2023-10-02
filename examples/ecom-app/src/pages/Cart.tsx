import { Link } from "react-router-dom";
import CryptoPay from "@sijibomi/cryptopay-sdk";

const Cart = () => {

    function handleCheckOut(e: React.MouseEvent<HTMLAnchorElement, MouseEvent>){
        e.preventDefault()
        const cryptoPay = new CryptoPay({
            price: 0.1,
            fiat: "USD",
            identifier: "sijibomiolajubu@gmail.com",
            apiUrl: "http://localhost:7000",
            apiKey: "bfdd8047-4f51-4965-9007-ff6500c0cfcd",
            currencies: ["btc"],
            onSuccess: function(voucher: any) {
                console.log("Successfully completed the payment.", voucher);
                return null
            },
            button: document.getElementById("checkout-cart-btn") as HTMLElement,
        })
        console.log(cryptoPay)
        cryptoPay.init()
    }

    return (
        <div className="max-w-6xl mx-auto">
            <div className="px-3 md:px-6">
                <table className="w-full">
                    <thead className="mb-4 cart-border-thick py-3">
                        <tr className="grid grid-cols-3 md:grid-cols-12">
                            <th className="md:col-span-8 text-left py-4">Product</th>
                            <th className="md:col-span-1 text-left py-4">Quantity</th>
                            <th className="md:col-span-1 text-left py-4">Price</th>
                            <th className="md:col-span-1 text-left py-4">Total</th>
                        </tr>
                    </thead>
                    <tbody className="py-3">
                        <tr className="grid grid-cols-3 md:grid-cols-12 py-2">
                            <td className="md:col-span-8 flex ">
                                <div className="flex  justify-center">
                                    <div className="">
                                        <figure>
                                            <img className="h-[100px] w-[90px]" src="https://images.unsplash.com/photo-1574180566232-aaad1b5b8450?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxzZWFyY2h8NTZ8fHQlMjBzaGlydHxlbnwwfHwwfHx8MA%3D%3D&auto=format&fit=crop&w=800&q=60" alt="shirt-1"/>
                                        </figure>
                                    </div>
                                    <div className="ml-5">
                                        <h4 className="text-lg font-bold">White Shirt packg I</h4>
                                        <h5 className="text-md">Status: <span className="text-green-600">In stock</span></h5>
                                    </div>
                                </div>
                            </td>
                            <td className="md:col-span-1 py-2">2</td>
                            <td className="md:col-span-1 py-2">$ 10.00</td>
                            <td className="md:col-span-1 py-2">$ 20.00</td>
                            <td className="md:col-span-1 py-2">
                                <div className="bg-red-500 p-1 rounded-[5px] py-0 rem-btn flex items-center justify-center">
                                    <svg xmlns="http://www.w3.org/2000/svg" width="12" height="12" fill="white" className="bi bi-trash" viewBox="0 0 16 16"> <path d="M5.5 5.5A.5.5 0 0 1 6 6v6a.5.5 0 0 1-1 0V6a.5.5 0 0 1 .5-.5zm2.5 0a.5.5 0 0 1 .5.5v6a.5.5 0 0 1-1 0V6a.5.5 0 0 1 .5-.5zm3 .5a.5.5 0 0 0-1 0v6a.5.5 0 0 0 1 0V6z"/> <path fill-rule="evenodd" d="M14.5 3a1 1 0 0 1-1 1H13v9a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V4h-.5a1 1 0 0 1-1-1V2a1 1 0 0 1 1-1H6a1 1 0 0 1 1-1h2a1 1 0 0 1 1 1h3.5a1 1 0 0 1 1 1v1zM4.118 4 4 4.059V13a1 1 0 0 0 1 1h6a1 1 0 0 0 1-1V4.059L11.882 4H4.118zM2.5 3V2h11v1h-11z"/> </svg>
                                    <Link className="text-sm text-white ml-2" to='/checkout'>Remove</Link>
                                </div>
                            </td>
                        </tr>

                        <tr className="grid grid-cols-3 md:grid-cols-12 py-2">
                            <td className="md:col-span-8 flex ">
                                <div className="flex  justify-center">
                                    <div className="">
                                        <figure>
                                            <img className="h-[100px] w-[90px]" src="https://images.unsplash.com/photo-1574180566232-aaad1b5b8450?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxzZWFyY2h8NTZ8fHQlMjBzaGlydHxlbnwwfHwwfHx8MA%3D%3D&auto=format&fit=crop&w=800&q=60" alt="shirt-1"/>
                                        </figure>
                                    </div>
                                    <div className="ml-5">
                                        <h4 className="text-lg font-bold">White Shirt packg I</h4>
                                        <h5 className="text-md">Status: <span className="text-green-600">In stock</span></h5>
                                    </div>
                                </div>
                            </td>
                            <td className="md:col-span-1 py-2">2</td>
                            <td className="md:col-span-1 py-2">$ 10.00</td>
                            <td className="md:col-span-1 py-2">$ 20.00</td>
                            <td className="md:col-span-1 py-2">
                                <div className="bg-red-500 p-1 rounded-[5px] py-0 rem-btn flex items-center justify-center">
                                    <svg xmlns="http://www.w3.org/2000/svg" width="12" height="12" fill="white" className="bi bi-trash" viewBox="0 0 16 16"> <path d="M5.5 5.5A.5.5 0 0 1 6 6v6a.5.5 0 0 1-1 0V6a.5.5 0 0 1 .5-.5zm2.5 0a.5.5 0 0 1 .5.5v6a.5.5 0 0 1-1 0V6a.5.5 0 0 1 .5-.5zm3 .5a.5.5 0 0 0-1 0v6a.5.5 0 0 0 1 0V6z"/> <path fill-rule="evenodd" d="M14.5 3a1 1 0 0 1-1 1H13v9a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V4h-.5a1 1 0 0 1-1-1V2a1 1 0 0 1 1-1H6a1 1 0 0 1 1-1h2a1 1 0 0 1 1 1h3.5a1 1 0 0 1 1 1v1zM4.118 4 4 4.059V13a1 1 0 0 0 1 1h6a1 1 0 0 0 1-1V4.059L11.882 4H4.118zM2.5 3V2h11v1h-11z"/> </svg>
                                    <Link className="text-sm text-white ml-2" to='/checkout'>Remove</Link>
                                </div>
                            </td>
                        </tr> 
                    </tbody>
                </table>

                <div className="flex flex-col cart-products pb-6">
                    <div className="text-right flex self-end justify-center items-center mb-0">
                        <h4 className="text-2xl font-semibold mr-4">Items Subtotal: </h4>
                        <h5 className="text-lg">$ 20.00</h5>
                    </div>

                    <div className="text-right flex self-end justify-center items-center">
                        <h4 className="text-2xl font-semibold mr-4">Shipping & Handling: </h4>
                        <h5 className="text-lg">$ 2.00</h5>
                    </div>
                </div>

                <div className="flex flex-col pb-6">
                    <div className="text-right flex self-end justify-center items-center mb-0">
                        <h4 className="text-2xl font-semibold mr-4">Total </h4>
                        <h5 className="text-lg">$ 20.00</h5>
                    </div>

                    <div className="text-right flex self-end justify-center items-center mt-3 ">
                        <Link className="text-sm ml-2 flex items-center cart-border p-2" to='/checkout'>
                            <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" className="bi bi-cart-fill" viewBox="0 0 16 16"> <path d="M0 1.5A.5.5 0 0 1 .5 1H2a.5.5 0 0 1 .485.379L2.89 3H14.5a.5.5 0 0 1 .491.592l-1.5 8A.5.5 0 0 1 13 12H4a.5.5 0 0 1-.491-.408L2.01 3.607 1.61 2H.5a.5.5 0 0 1-.5-.5zM5 12a2 2 0 1 0 0 4 2 2 0 0 0 0-4zm7 0a2 2 0 1 0 0 4 2 2 0 0 0 0-4zm-7 1a1 1 0 1 1 0 2 1 1 0 0 1 0-2zm7 0a1 1 0 1 1 0 2 1 1 0 0 1 0-2z"/> </svg>
                            <h4 className="ml-2">Continue shopping</h4>
                        </Link>
                        <Link className="text-sm ml-2 flex items-center rounded-[5px] p-2 bg-green-500" id="checkout-cart-btn" to="#" onClick={(e) => handleCheckOut(e)}>
                            {/* "please wait while" loading */}
                            <h4>Checkout</h4>
                        </Link>
                    </div>
                </div>
            </div>
        </div>
    )
}

export default Cart;