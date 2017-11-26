import Foundation

public protocol ILocationUseCase {
    func getLocations()
}

class LocationUseCase: ILocationUseCase {
    
    let repository: ILocationRepository = {
        return LocationRepository()
    }()
    
    func getLocations() {
        repository.getLocations()
    }
}
