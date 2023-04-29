import { FormEvent, RefObject, useRef } from "react"
import Main from "../layouts/main";
import { AppDispatch } from "@/app/store";
import { useDispatch } from "react-redux";
import { userRegister } from "@/features/user/userSlice";
import { RegisterBody } from "@/models/request";
import { resultEntity } from "@/models/result";
import { useRouter } from "next/router";

const Resister = () => {
    const router = useRouter();
    const dispatch:AppDispatch = useDispatch()

    const textEmail = useRef<HTMLInputElement>(null)
    const textPassword = useRef<HTMLInputElement>(null)
    const textName = useRef<HTMLInputElement>(null)

    const register = async (e: FormEvent<HTMLFormElement>) => {
        e.preventDefault()
        // e.stopPropagation()

        const body:RegisterBody = {
            email: textEmail.current?.value,
            password: textPassword.current?.value,
            name: textName.current?.value
        }

        const res = await dispatch(userRegister(body))
        .then((res) => {

            if(res.meta.requestStatus === 'fulfilled') {
                const _data = res.payload as {message: string}
    
                console.table(_data.message)

                return {
                    message: 'OK'
                } as resultEntity
            }
            
            console.log(res.payload as number)

            return {
                message: res.payload as number
            } as resultEntity
        })

        if(res.message == 'OK') {
            router.push('/login');
        }
    }

    return (
        <Main>
            <div className="flex justify-center items-center w-full md:w-[500px] h-[500px] mx-auto">            
                <RegisterForm textEmail={textEmail} textPassword={textPassword} textName={textName} register={register} />
            </div>
        </Main>
    )
}

const RegisterForm = ({
        textEmail, 
        textPassword, 
        textName,
        register
    }:{
        textEmail: RefObject<HTMLInputElement>,
        textPassword: RefObject<HTMLInputElement>,
        textName: RefObject<HTMLInputElement>,
        register: (e: FormEvent<HTMLFormElement>) => void
    }) => {
    return (
        <div className="p-8 w-full h-full bg-slate-600 flex flex-col rounded-2xl">
             <h2 className="text-2xl mb-2">Register</h2>
             <hr className="mb-4" />
            <form className="p-4 pt-0" onSubmit={register}>
                <div className="mb-8">
                    <div className="mb-4">
                        <div className="mb-2">Email:</div>
                        <input 
                            ref={textEmail}
                            type="email" 
                            required
                            className="w-full p-2 text-black outline-0" 
                        />
                    </div>   
                    <div className="mb-4">
                        <div className="mb-2">Password:</div>
                        <input 
                            ref={textPassword}
                            type="password"                     
                            required
                            className="w-full p-2 text-black outline-0"
                        />
                    </div> 
                    <div className="mb-4">
                        <div className="mb-2">Name:</div>
                        <input 
                            ref={textName}
                            type="text"                     
                            required
                            className="w-full p-2 text-black outline-0"
                        />
                    </div>
                </div>
                <div>
                    <button 
                        role="button"
                        type="submit"
                        className="w-1/5 p-1 rounded-lg bg-red-800"                                             
                    >
                        Register
                    </button>
                </div>
            </form>
        </div>
    )
}

export default Resister