import "./Navbar.scss"

export const Navbar = ({ home }) => {
    return (
        <div className="navbar">
            <div className="navbar-header">
                <h1 className="navbar-logo">
                    {home ? (
                        <a href="/about">
                            DNA matcher
                        </a>
                    ) : (
                        <a href="/">
                            TestDNA
                        </a>
                    )}
                </h1>
            </div>
        </div>
    )
}