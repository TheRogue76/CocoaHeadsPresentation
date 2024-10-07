//
//  File.swift
//  
//
//  Created by Parsa's Content Creation Corner on 2024-10-07.
//

import Foundation
import GRPC
import NIO

class FoodCompService {
    private let group = PlatformSupport.makeEventLoopGroup(loopCount: 1)
    static let shared = FoodCompService()
    private let client: FoodCompanyGateWayServiceClientProtocol
    private init() {
        let channel = ClientConnection(configuration: .default(target: ConnectionTarget.hostAndPort("127.0.0.1", 50051), eventLoopGroup: group))
        client = FoodCompanyGateWayServiceNIOClient(channel: channel)
    }
    
    func getOrderStatus() {
        var request = GetOrderStatusRequest()
        request.id = ""
        let calloption = CallOptions(eventLoopPreference: .indifferent)
            // Make the RPC call to the server.
        let unaryCall = client.getOrderStatus(request, callOptions: calloption)
            // wait() on the response to stop the program from exiting before the response is received.
        do {
            let response = try unaryCall.response.wait()
            print("Sum Received received: \(response)")
        } catch {
            print("Sum Received failed: \(error)")
            return
        }
    }
    
    func getOrderStatusAsync() async -> Void {
        var request = GetOrderStatusRequest()
        request.id = ""
        let calloption = CallOptions(eventLoopPreference: .indifferent)
            // Make the RPC call to the server.
        let unaryCall = client.getOrderStatus(request, callOptions: calloption)
            // wait() on the response to stop the program from exiting before the response is received.
        do {
            let response = try await unaryCall.response.get()
            print("Sum Received received: \(response)")
        } catch {
            print("Sum Received failed: \(error)")
            return
        }
    }
}

