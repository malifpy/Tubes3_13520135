import { Navbar } from "../component/Navbar/Navbar"
import "./About.scss"

export const About = () => {
    return (
        <>
            <Navbar />
            <div className="about-wrapper">
                <p>Made by</p>
                <div className="about-container">
                    <div className="about-text">
                        <ul>Muhammad Alif Putra Yasa</ul>
                        <ul>13520135</ul>
                        <ul>K03</ul>
                        <ul>String Matching and Regular Expression Algorithm</ul>
                    </div>
                </div>
                <div className="about-container-reverse">
                    <div className="about-text">
                        <ul>Azmi Alfatih Shalahuddin</ul>
                        <ul>13520158</ul>
                        <ul>K02</ul>
                    </div>
                </div>
                <div className="about-container">
                    <div className="about-text">
                        <ul>Ghazian Tsabit Alkamil</ul>
                        <ul>13520165</ul>
                        <ul>K03</ul>
                        <ul>FE and BE Handler</ul>
                    </div>
                </div>
            </div>
        </>
    )
}