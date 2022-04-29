import AddPenyakit from "../component/AddPenyakit"
import TestDNA from "../component/TestDNA"
import { Navbar } from "../component/Navbar/Navbar"
import HasilPrediksi from "../component/HasilPrediksi"
import "./MatcherDNA.scss"


export const MatcherDNA = () => {
    return (
        <>
            <Navbar 
                home
            />
            <div className="compressor-wrapper">
                <AddPenyakit />
                <TestDNA />  
                <HasilPrediksi />
            </div>
        </>
    )
}