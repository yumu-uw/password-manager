import { InfoOutlineIcon } from "@chakra-ui/icons";
import {
	Box,
	Flex,
	IconButton,
	Spacer,
	StackDivider,
	VStack,
} from "@chakra-ui/react";
import { useAtomValue } from "jotai";
import { Suspense } from "react";
import { GetPlacePasswords } from "../../wailsjs/go/main/App";
import placeInfoAtom from "../atoms/placeInfoAtom";

function PlaceTable() {
	const placeInfo = useAtomValue(placeInfoAtom);

	const showDetailModal = (id: number) => {
		GetPlacePasswords(id).then((result) => console.log(result));
	};

	return (
		<Suspense fallback="Loading...">
			<VStack
				divider={<StackDivider borderColor="gray.200" />}
				spacing={1}
				align="stretch"
			>
				{placeInfo.map((place) => (
					<Flex key={place.PlaceName}>
						<Box bg="red.400">{place.PlaceName}</Box>
						<Spacer />
						<Box>
							<IconButton
								aria-label="Search database"
								icon={<InfoOutlineIcon />}
								onClick={() => showDetailModal(place.ID)}
							/>
						</Box>
					</Flex>
				))}
			</VStack>
		</Suspense>
	);
}

export default PlaceTable;
