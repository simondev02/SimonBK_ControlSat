package models

var Consultas = map[int]string{

	// 2    "Finanzauto"
	// 3    "Motorysa Demo"
	// 4    "Carfiao"
	// 11    "PrestAuto"
	// 12    "Banco Finandina"
	// 19    "Equirent

	// Finandina
	12: `SELECT Imei, Plate, Description,Speed,Latitude, Longitude, Timestamp, Event, Odometer/1000 FROM (
        SELECT	Devices.Imei, 
                Vehicles.Plate, 
                SubFleets.Description,
                LastsPositions.Latitude,
                LastsPositions.Longitude,
				LastsPositions.Speed,
                LastsPositions.Timestamp,
                LastsPositions.Event,
				LastsPositions.Odometer,
                ROW_NUMBER() OVER (
                    partition by Imei, Plate
                    Order by Timestamp DESC
                ) as rn 
        FROM    [VisualSat.Avl.Database_finandina].[dbo].[Devices] WITH (NOLOCK) INNER JOIN
                [VisualSat.Avl.Database_finandina].[dbo].DevicesVehiclesOdometer WITH (NOLOCK) ON Devices.Id = DevicesVehiclesOdometer.DeviceId INNER JOIN
                [VisualSat.Avl.Database_finandina].[dbo].Vehicles WITH (NOLOCK) ON DevicesVehiclesOdometer.VehicleId = Vehicles.Id INNER JOIN
                [VisualSat.Avl.Database_finandina].[dbo].LastsPositions WITH(NOLOCK) ON Devices.Device_Id = LastsPositions.Device_Id INNER JOIN
                [VisualSat.Avl.Database_finandina].[dbo].SubFleets WITH (NOLOCK) ON Vehicles.Subfleet_Id = SubFleets.Id
        WHERE   Devices.Imei IN ({IMEI})
                    GROUP BY Devices.Imei, Vehicles.Plate, SubFleets.Description,LastsPositions.Speed,LastsPositions.Latitude,LastsPositions.Longitude, LastsPositions.Timestamp, LastsPositions.Event, LastsPositions.Odometer ) AS SQ_1 
                    WHERE rn = 1 ORDER BY Timestamp desc`, // Consula Finandina
	// Finanzauto
	2: `SELECT Imei, Plate, Description,Speed,Latitude, Longitude, Timestamp, Event, Odometer/1000 FROM (
		SELECT	Devices.Imei, 
				Vehicles.Plate, 
				SubFleets.Description,
				LastsPositions.Latitude,
				LastsPositions.Longitude,
				LastsPositions.Speed,
				LastsPositions.Timestamp,
				LastsPositions.Event,
				LastsPositions.Odometer,
				ROW_NUMBER() OVER (
					partition by Imei, Plate
					Order by Timestamp DESC
				) as rn 
		FROM    [VisualSat.Avl.Database_col].[dbo].[Devices] WITH (NOLOCK) INNER JOIN
				[VisualSat.Avl.Database_col].[dbo].DevicesVehiclesOdometer WITH (NOLOCK) ON Devices.Id = DevicesVehiclesOdometer.DeviceId INNER JOIN
				[VisualSat.Avl.Database_col].[dbo].Vehicles WITH (NOLOCK) ON DevicesVehiclesOdometer.VehicleId = Vehicles.Id INNER JOIN
				[VisualSat.Avl.Database_col].[dbo].LastsPositions WITH(NOLOCK) ON Devices.Device_Id = LastsPositions.Device_Id INNER JOIN
				[VisualSat.Avl.Database_col].[dbo].SubFleets WITH (NOLOCK) ON Vehicles.Subfleet_Id = SubFleets.Id
		WHERE   Devices.Imei IN ({IMEI})
							GROUP BY Devices.Imei, Vehicles.Plate, SubFleets.Description,LastsPositions.Speed,LastsPositions.Latitude,LastsPositions.Longitude, LastsPositions.Timestamp, LastsPositions.Event, LastsPositions.Odometer ) AS SQ_1 
							WHERE rn = 1 ORDER BY Timestamp desc
		`, // Consulta Finanzauto

	//Carfiao
	4: `SELECT Imei, Plate, Description,Speed ,Latitude, Longitude, Timestamp, Event, Odometer/1000 FROM (
		SELECT	Devices.Imei, 
				Vehicles.Plate, 
				SubFleets.Description,
				LastsPositions.Latitude,
				LastsPositions.Longitude,
				LastsPositions.Speed,
				LastsPositions.Timestamp,
				LastsPositions.Event,
				LastsPositions.Odometer,
				ROW_NUMBER() OVER (
					partition by Imei, Plate
					Order by Timestamp DESC
				) as rn 
		FROM    [VisualSat.Avl.Database_col].[dbo].[Devices] WITH (NOLOCK) INNER JOIN
				[VisualSat.Avl.Database_col].[dbo].DevicesVehiclesOdometer WITH (NOLOCK) ON Devices.Id = DevicesVehiclesOdometer.DeviceId INNER JOIN
				[VisualSat.Avl.Database_col].[dbo].Vehicles WITH (NOLOCK) ON DevicesVehiclesOdometer.VehicleId = Vehicles.Id INNER JOIN
				[VisualSat.Avl.Database_col].[dbo].LastsPositions WITH(NOLOCK) ON Devices.Device_Id = LastsPositions.Device_Id INNER JOIN
				[VisualSat.Avl.Database_col].[dbo].SubFleets WITH (NOLOCK) ON Vehicles.Subfleet_Id = SubFleets.Id
		WHERE        	Devices.Imei IN ({IMEI})
							GROUP BY Devices.Imei, Vehicles.Plate, SubFleets.Description,LastsPositions.Speed,LastsPositions.Latitude,LastsPositions.Longitude, LastsPositions.Timestamp, LastsPositions.Event, LastsPositions.Odometer, ) AS SQ_1 
							WHERE rn = 1 ORDER BY Timestamp desc
		`, // Consulta de Carfiao
	//PrestaAuto
	11: `SELECT Imei, Plate, Description,Speed,Latitude, Longitude, Timestamp, Event, Odometer/1000 FROM (
		SELECT	Devices.Imei, 
				Vehicles.Plate, 
				SubFleets.Description,
				LastsPositions.Latitude,
				LastsPositions.Longitude,
				LastsPositions.Speed,
				LastsPositions.Timestamp,
				LastsPositions.Event,
				LastsPositions.Odometer,
				ROW_NUMBER() OVER (
					partition by Imei, Plate
					Order by Timestamp DESC
				) as rn 
		FROM    [VisualSat.Avl.Database_finandina].[dbo].[Devices] WITH (NOLOCK) INNER JOIN
				[VisualSat.Avl.Database_finandina].[dbo].DevicesVehiclesOdometer WITH (NOLOCK) ON Devices.Id = DevicesVehiclesOdometer.DeviceId INNER JOIN
				[VisualSat.Avl.Database_finandina].[dbo].Vehicles WITH (NOLOCK) ON DevicesVehiclesOdometer.VehicleId = Vehicles.Id INNER JOIN
				[VisualSat.Avl.Database_finandina].[dbo].LastsPositions WITH(NOLOCK) ON Devices.Device_Id = LastsPositions.Device_Id INNER JOIN
				[VisualSat.Avl.Database_finandina].[dbo].SubFleets WITH (NOLOCK) ON Vehicles.Subfleet_Id = SubFleets.Id
		WHERE        	Devices.Imei IN ({IMEI})
							GROUP BY Devices.Imei, Vehicles.Plate, SubFleets.Description,LastsPositions.Speed,LastsPositions.Latitude,LastsPositions.Longitude, LastsPositions.Timestamp, LastsPositions.Event, LastsPositions.Odometer ) AS SQ_1 
							WHERE rn = 1 ORDER BY Timestamp desc`, // Consulta de Prestautos
}
