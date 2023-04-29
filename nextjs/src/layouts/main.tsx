import Footer from "@/components/footer"
import Header from "@/components/header"
import { ReactNode } from "react"

const Main = ({children}: {children: ReactNode}) => {
    return (
        <>
            <Header/>
            {children}
            <Footer/>
        </>
    )
}

export default Main