// Gola Score Keeper - Main JavaScript

class GolaScoreKeeper {
    constructor() {
        this.currentContestant = null;
        this.currentScore = 0;
        this.attemptNumber = 1;
        this.websocket = null;
        this.cameraStream = null;
        
        this.initializeEventListeners();
        this.initializeWebSocket();
        this.loadLeaderboard();
    }

    initializeEventListeners() {
        // Badge scanning
        document.getElementById('scanButton').addEventListener('click', () => this.startBadgeScanning());
        
        // Contestant actions
        document.getElementById('startScoring').addEventListener('click', () => this.startScoring());
        document.getElementById('clearContestant').addEventListener('click', () => this.clearContestant());
        
        // Scoring buttons
        document.getElementById('goalButton').addEventListener('click', () => this.recordShot('goal'));
        document.getElementById('missButton').addEventListener('click', () => this.recordShot('miss'));
        
        // Scoring actions
        document.getElementById('undoLastShot').addEventListener('click', () => this.undoLastShot());
        document.getElementById('finishScoring').addEventListener('click', () => this.finishScoring());
        
        // Leaderboard
        document.getElementById('refreshLeaderboard').addEventListener('click', () => this.loadLeaderboard());
        document.getElementById('leaderboardFilter').addEventListener('change', () => this.loadLeaderboard());
    }

    initializeWebSocket() {
        const protocol = window.location.protocol === 'https:' ? 'wss:' : 'ws:';
        const wsUrl = `${protocol}//${window.location.host}/ws/leaderboard`;
        
        this.websocket = new WebSocket(wsUrl);
        
        this.websocket.onopen = () => {
            console.log('WebSocket connected');
        };
        
        this.websocket.onmessage = (event) => {
            const data = JSON.parse(event.data);
            this.handleWebSocketMessage(data);
        };
        
        this.websocket.onclose = () => {
            console.log('WebSocket disconnected');
            // Attempt to reconnect after 5 seconds
            setTimeout(() => this.initializeWebSocket(), 5000);
        };
        
        this.websocket.onerror = (error) => {
            console.error('WebSocket error:', error);
        };
    }

    handleWebSocketMessage(data) {
        switch (data.type) {
            case 'leaderboard_update':
                this.updateLeaderboard(data.leaderboard);
                break;
            case 'score_update':
                this.updateScoreDisplay(data.score);
                break;
            default:
                console.log('Unknown WebSocket message type:', data.type);
        }
    }

    async startBadgeScanning() {
        const scanButton = document.getElementById('scanButton');
        const scanResult = document.getElementById('scanResult');
        
        try {
            scanButton.textContent = 'Scanning...';
            scanButton.disabled = true;
            
            // Start camera if not already started
            if (!this.cameraStream) {
                await this.startCamera();
            }
            
            // Simulate badge scanning (replace with actual OCR implementation)
            const result = await this.simulateBadgeScan();
            
            if (result.success) {
                this.currentContestant = result.contestant;
                this.displayContestantInfo(result.contestant);
                scanResult.innerHTML = '<div class="success">Badge scanned successfully!</div>';
            } else {
                scanResult.innerHTML = '<div class="error">Failed to scan badge. Please try again.</div>';
            }
            
        } catch (error) {
            console.error('Badge scanning error:', error);
            scanResult.innerHTML = '<div class="error">Error scanning badge. Please try again.</div>';
        } finally {
            scanButton.textContent = 'Scan Badge';
            scanButton.disabled = false;
        }
    }

    async startCamera() {
        try {
            this.cameraStream = await navigator.mediaDevices.getUserMedia({
                video: { 
                    width: 1280, 
                    height: 720,
                    facingMode: 'environment' // Use back camera if available
                }
            });
            
            const video = document.getElementById('cameraPreview');
            video.srcObject = this.cameraStream;
            
        } catch (error) {
            console.error('Camera access error:', error);
            throw new Error('Unable to access camera. Please check permissions.');
        }
    }

    async simulateBadgeScan() {
        // Simulate OCR processing delay
        await new Promise(resolve => setTimeout(resolve, 2000));
        
        // Simulate successful scan with mock data
        return {
            success: true,
            contestant: {
                id: 1,
                name: 'John Doe',
                company: 'Tech Corp',
                email: 'john.doe@techcorp.com',
                badge_id: 'BADGE123'
            }
        };
    }

    displayContestantInfo(contestant) {
        document.getElementById('contestantName').textContent = contestant.name;
        document.getElementById('contestantCompany').textContent = contestant.company;
        document.getElementById('contestantEmail').textContent = contestant.email;
        
        document.getElementById('contestantInfo').style.display = 'block';
        document.getElementById('contestantInfo').classList.add('fade-in');
    }

    startScoring() {
        this.currentScore = 0;
        this.attemptNumber = 1;
        
        document.getElementById('scoringInterface').style.display = 'block';
        document.getElementById('scoringInterface').classList.add('fade-in');
        
        this.updateScoreDisplay();
    }

    async recordShot(result) {
        try {
            const shotData = {
                contestant_id: this.currentContestant.id,
                shot_result: result,
                attempt_number: this.attemptNumber,
                operator_id: 'operator-001',
                session_id: 'session-' + Date.now()
            };
            
            const response = await fetch('/api/penalty-shots', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(shotData)
            });
            
            if (response.ok) {
                if (result === 'goal') {
                    this.currentScore++;
                }
                this.attemptNumber++;
                this.updateScoreDisplay();
                
                // Show visual feedback
                this.showShotFeedback(result);
                
            } else {
                throw new Error('Failed to record shot');
            }
            
        } catch (error) {
            console.error('Error recording shot:', error);
            alert('Failed to record shot. Please try again.');
        }
    }

    showShotFeedback(result) {
        const feedback = document.createElement('div');
        feedback.className = `shot-feedback ${result}`;
        feedback.textContent = result.toUpperCase();
        
        document.body.appendChild(feedback);
        
        setTimeout(() => {
            feedback.remove();
        }, 2000);
    }

    updateScoreDisplay() {
        document.getElementById('attemptNumber').textContent = this.attemptNumber;
        document.getElementById('currentScore').textContent = this.currentScore;
    }

    async undoLastShot() {
        if (this.attemptNumber > 1) {
            this.attemptNumber--;
            if (this.currentScore > 0) {
                this.currentScore--;
            }
            this.updateScoreDisplay();
            
            // TODO: Implement actual undo functionality via API
            console.log('Undo last shot');
        }
    }

    async finishScoring() {
        try {
            // TODO: Implement finish scoring functionality
            console.log('Finishing scoring session');
            
            // Hide scoring interface
            document.getElementById('scoringInterface').style.display = 'none';
            
            // Clear contestant
            this.clearContestant();
            
            // Refresh leaderboard
            this.loadLeaderboard();
            
        } catch (error) {
            console.error('Error finishing scoring:', error);
        }
    }

    clearContestant() {
        this.currentContestant = null;
        this.currentScore = 0;
        this.attemptNumber = 1;
        
        document.getElementById('contestantInfo').style.display = 'none';
        document.getElementById('scoringInterface').style.display = 'none';
        
        // Stop camera
        if (this.cameraStream) {
            this.cameraStream.getTracks().forEach(track => track.stop());
            this.cameraStream = null;
        }
    }

    async loadLeaderboard() {
        try {
            const filter = document.getElementById('leaderboardFilter').value;
            const response = await fetch(`/api/leaderboard?filter=${filter}`);
            
            if (response.ok) {
                const leaderboard = await response.json();
                this.updateLeaderboard(leaderboard);
            } else {
                throw new Error('Failed to load leaderboard');
            }
            
        } catch (error) {
            console.error('Error loading leaderboard:', error);
            // Show mock data for development
            this.showMockLeaderboard();
        }
    }

    updateLeaderboard(leaderboard) {
        const tbody = document.getElementById('leaderboardBody');
        tbody.innerHTML = '';
        
        leaderboard.forEach((contestant, index) => {
            const row = document.createElement('tr');
            row.className = `rank-${index + 1}`;
            
            row.innerHTML = `
                <td>${index + 1}</td>
                <td>${contestant.name}</td>
                <td>${contestant.company || 'N/A'}</td>
                <td>${contestant.successful_shots}</td>
                <td>${contestant.total_attempts}</td>
                <td>${(contestant.success_percentage * 100).toFixed(1)}%</td>
            `;
            
            tbody.appendChild(row);
        });
    }

    showMockLeaderboard() {
        const mockData = [
            { name: 'John Doe', company: 'Tech Corp', successful_shots: 8, total_attempts: 10, success_percentage: 0.8 },
            { name: 'Jane Smith', company: 'Innovation Inc', successful_shots: 7, total_attempts: 10, success_percentage: 0.7 },
            { name: 'Mike Johnson', company: 'Startup Co', successful_shots: 6, total_attempts: 8, success_percentage: 0.75 },
            { name: 'Sarah Wilson', company: 'Tech Solutions', successful_shots: 5, total_attempts: 7, success_percentage: 0.714 },
            { name: 'David Brown', company: 'Digital Corp', successful_shots: 4, total_attempts: 6, success_percentage: 0.667 }
        ];
        
        this.updateLeaderboard(mockData);
    }
}

// Initialize application when DOM is loaded
document.addEventListener('DOMContentLoaded', () => {
    new GolaScoreKeeper();
});

// Add CSS for shot feedback
const style = document.createElement('style');
style.textContent = `
    .shot-feedback {
        position: fixed;
        top: 50%;
        left: 50%;
        transform: translate(-50%, -50%);
        font-size: 4rem;
        font-weight: bold;
        color: white;
        padding: 20px 40px;
        border-radius: 10px;
        z-index: 1000;
        animation: feedbackPulse 2s ease-out;
    }
    
    .shot-feedback.goal {
        background-color: #27ae60;
    }
    
    .shot-feedback.miss {
        background-color: #e74c3c;
    }
    
    @keyframes feedbackPulse {
        0% { transform: translate(-50%, -50%) scale(0.5); opacity: 0; }
        50% { transform: translate(-50%, -50%) scale(1.2); opacity: 1; }
        100% { transform: translate(-50%, -50%) scale(1); opacity: 0; }
    }
`;
document.head.appendChild(style);
