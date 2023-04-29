import { AppDispatch } from "@/app/store"
import { userLogin, userProfile } from "@/features/user/userSlice"
import Main from "@/layouts/main"
import { useRouter } from "next/router"
import { RefObject, useRef, useState } from "react"
import { useDispatch } from "react-redux"

const Login = () => {
    const router = useRouter();
    const dispatch:AppDispatch = useDispatch()

    const [message, setMessage] = useState<string>("")
    
    const textEmail = useRef<HTMLInputElement>(null)
    const textPassword = useRef<HTMLInputElement>(null)

    const login = async (email: string, password: string) => {
    
        const res = await dispatch(userLogin({
            email: email,
            password: password
        }))
        .then((res) => {
            if(res.meta.requestStatus === 'fulfilled') {
                const _data = res.payload as {accessToken: string}
    
                console.table(_data.accessToken)

                return true
            }
            
            console.log(res.payload as number)
            return false
        })

        if(res) {
            // Call api get user profile
            console.log("Call api get user profile")

            await dispatch(userProfile())
            .then((res) => {
                console.table(res)

                if(res.meta.requestStatus === 'fulfilled') {
                    router.push('/')
                }
            })
        }else {
            setMessage("Email or Password incorrect")
        }
        
    }

    return (
        <Main>
            <div className="flex justify-center flex-col items-center w-full md:w-[500px] mx-auto">
                <div className="flex justify-center text-red-400 text-lg mb-3">
                    {message}
                </div>
                <LoginForm textEmail={textEmail} textPassword={textPassword} login={login} />
            </div>
        </Main>
    )
}

const LoginForm = ({
        textEmail,
        textPassword,
        login
    }: {
        textEmail: RefObject<HTMLInputElement>,
        textPassword: RefObject<HTMLInputElement>
        login: (email: string, password: string) => void
    }) => {
    return (
        <form className="p-8 w-full h-full bg-slate-600 flex flex-col rounded-2xl" onSubmit={(e) => {
            e.preventDefault()
            login(textEmail.current?.value!, textPassword.current?.value!)
        }}>
            <h2 className="text-2xl mb-2">Login</h2>
            <hr className="mb-4" />
            <div className="mb-4">
                <input 
                    ref={textEmail}
                    type="email" 
                    placeholder="Email"
                    required
                    className="w-full p-2 text-black outline-0" 
                />
            </div>
            <div className="mb-4">
                <input 
                    ref={textPassword}
                    type="password" 
                    placeholder="Password"
                    required
                    className="w-full p-2 text-black outline-0"
                />
            </div>
            <div>
                <button 
                    role="button"
                    type="submit"
                    className="w-1/5 p-1 rounded-lg bg-red-800"                                             
                >
                    Login
                </button>
            </div>
        </form>
    )
}

export default Login