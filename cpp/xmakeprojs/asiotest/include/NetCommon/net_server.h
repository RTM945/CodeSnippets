#pragma once

#include "net_common.h"
#include "net_tsqueue.h"
#include "net_message.h"
#include "net_connection.h"

namespace olc
{
    namespace net
    {
        template<typename  T>
        class server_interface
        {
        public:
            server_interface(uint16_t port)
                : m_asioAccptor(m_asioContext, asio::ip::tcp::endpoint(asio::ip::tcp::v4(), port))
            {

            }

            virtual ~server_interface()
            {
                Stop();
            }

            bool Start()
            {
                try {
                    WaitForClientConnection();

                    m_threadContext = std::thread([this]() {m_asioContext.run(); });
                } catch (std::exception& e) {
                    std::cerr << "[SERVER] Exception: " << e.what() << "\n";
                    return false;
                }
                std::cout << "[SERVER] Started!\n";
                return true;
            }

            void Stop()
            {
                m_asioContext.stop();
                if (m_threadContext.joinable()) 
                {
                    m_threadContext.join();
                }
                std::cout << "[SERVER] Stopped!\n";
            }

            // ASYNC
            void WaitForClientConnection()
            {
                m_asioAccptor.async_accept(
                    [this](std::error_code ec, asio::ip::tcp::socket socket) 
                    {
                        if (!ec) 
                        {
                            std::cout << "[SERVER] New Connection: " << socket.remote_endpoint() << "\n";

                            std::shared_ptr<connection<T>> newconn = 
                                std::make_shared<connection<T>>(
                                    connection<T>::owner::server, 
                                    m_asioContext, std::move(socket), m_qMessageIn
                                );
                            
                            if(OnClientConnect(newconn))
                            {
                                m_deqConnections.push_back(std::move(newconn));
                                m_deqConnections.back()->ConnectToClient(nIDCounter++);
                                
                                std::cout << "[" << m_deqConnections.back()->GetID() << "] Connection Approved\n";
                            }
                            else
                            {
                                std::cout << "[---------] Connection Denied\n";
                            }
                        }
                        else
                        {
                            std::cout << "[SERVER] New Connection Error: " << ec.message() << "\n";
                        }

                        WaitForClientConnection();
                    }

                );
            }

            // Send message to a specific client
            void MessageClient(std::shared_ptr<connection<T>> client, const message<T>& msg)
            {
                if (client && client->IsConnected())
                {
                    client->Send(msg);
                }
                else
                {
                    OnClientDisconnect(client);
                    client.reset();
                    m_deqConnections.erase(
                        std::remove(m_deqConnections.begin(), m_deqConnections.end(), client), 
                        m_deqConnections.end()
                    );
                }
            }

            // Send message to ALL clients
            void MessageAllClients(const message<T>& msg, std::shared_ptr<connection<T>> pIngoreClient = nullptr)
            {
                bool bInvaildClientExists = false;
                
                for (auto& client : m_deqConnections) 
                {
                    if (client && client->IsConnected())
                    {
                        if (client != pIngoreClient) 
                        {
                            client->Send(msg);
                        }
                    }
                    else 
                    {
                        OnClientDisconnect(client);
                        client.reset();
                        bInvaildClientExists = true;
                    }
                }

                m_deqConnections.erase(
                        std::remove(m_deqConnections.begin(), m_deqConnections.end(), nullptr), 
                        m_deqConnections.end()
                    );
            }

            void Update(size_t nMaxMessages = -1)
            {
                size_t nMessageCount = 0;
                while (nMessageCount < nMaxMessages && !m_qMessageIn.empty()) 
                {
                    auto msg = m_qMessageIn.pop_front();

                    OnMessage(msg.remote, msg.msg);

                    nMessageCount++;
                }
            }

        protected:
            virtual bool OnClientConnect(std::shared_ptr<connection<T>> client)
            {
                return false;
            }

            virtual void OnClientDisconnect(std::shared_ptr<connection<T>> client)
            {

            }

            virtual void OnMessage(std::shared_ptr<connection<T>> client, message<T>& msg)
            {

            }

        protected:
            tsqueue<owned_message<T>> m_qMessageIn;
            std::deque<std::shared_ptr<connection<T>>> m_deqConnections;

            asio::io_context m_asioContext;
            std::thread m_threadContext;

            asio::ip::tcp::acceptor m_asioAccptor;

            uint32_t nIDCounter = 10000;
        };
    }

}