import { Box, Button } from "@chakra-ui/react";
import { AddPassword, DBAccess } from "../wailsjs/go/main/App";
import PlaceTable from "./components/PlaceTable";

function App() {
	function dbaccess() {
		DBAccess().then((result) => console.log(result));
	}

	function addPassword() {
		AddPassword(5, "取引パスワード", "hogeHuGa").then((result) =>
			console.log(result),
		);
	}

	return (
		<Box height={"100vh"} bgColor={""}>
			<Box>
				<Button onClick={dbaccess}>DBAccess</Button>
				<Button onClick={addPassword}>Add</Button>
			</Box>
			<PlaceTable />
		</Box>
	);
}

export default App;
