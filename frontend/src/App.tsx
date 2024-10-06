import './App.css';
import { AddPassword, DBAccess } from "../wailsjs/go/main/App";

function App() {

    function dbaccess() {
        DBAccess().then((result) => console.log(result));
    }

    function addPassword() {
        AddPassword(5, "取引パスワード", "hogeHuGa").then((result) => console.log(result));
    }

    return (
        <div id="App">
            <div id="input" className="input-box">
                {/* biome-ignore lint/a11y/useButtonType: <explanation> */}
                <button className="btn" onClick={dbaccess}>DBAccess</button>
                {/* biome-ignore lint/a11y/useButtonType: <explanation> */}
                <button className="btn" onClick={addPassword}>Add</button>
            </div>
        </div>
    )
}

export default App
