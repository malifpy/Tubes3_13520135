import AddPenyakit from "../component/AddPenyakit"
import { Navbar } from "../component/Navbar/Navbar"

import "./TestDNA.scss"

export const TestDNA = () => {
    return (
        <>
            <Navbar 
                home
            />
            <div className="compressor-wrapper">
                <AddPenyakit />
            </div>
        </>
    )
}