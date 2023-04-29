import { useAppSelector } from "@/app/hooks"
import { RootState } from "@/app/store"
import Link from "next/link"

const Header = () => {

    // const accessToken = useAppSelector((state: RootState) => state.user.accessToken)
    const name = useAppSelector((state: RootState) => state.user.name)

    return (
        <header className="mb-12">
            <div className="w-full p-4 bg-green-700 mb-2">
                {
                    name ? (<h2 className="text-4xl text-center">Hello, {name}</h2>): ''
                }
            </div>
            <div className="flex justify-between w-[200px] px-3">
                <Link className="font-medium text-blue-600 dark:text-blue-500 hover:underline" href="/">Home</Link>
                <Link className="font-medium text-blue-600 dark:text-blue-500 hover:underline" href="/login">Login</Link>
                <Link className="font-medium text-blue-600 dark:text-blue-500 hover:underline" href="/register">Register</Link>
            </div>
            <hr className="w-3/4 mx-auto my-5" />
        </header>
    )
}

export default Header