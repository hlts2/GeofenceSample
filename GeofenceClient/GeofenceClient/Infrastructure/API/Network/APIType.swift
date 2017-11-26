import Foundation

struct APIType {
    struct GeofenceManagementService {
        struct Login: HTTPRequest {
            var baseURL: String {
                get {
                    return ""
                }
            }
            
            var path: String {
                get {
                    return ""
                }
            }
            
            var method: HTTPMethod
            
            var params: [String : String]
            
        }
        
        struct GetLocations: HTTPRequest {
            var baseURL: String {
                get {
                    return ""
                }
            }
            
            var path: String {
                get {
                    return ""
                }
            }
            
            var method: HTTPMethod
            
            var params: [String: String]
        }
    }
}
